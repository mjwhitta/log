package log

import (
	"os"
	"sync"

	hl "gitlab.com/mjwhitta/hilighter"
)

// CloseHandler is a function pointer. CloseHandlers are called when
// the Messengers is closed and allow for closing of files or sockets.
type CloseHandler func() error

// MsgHandler is a function pointer. MsgHandlers are called when a
// message is logged and allow for custom actions like writing to a
// file or a socket.
type MsgHandler func(msg *Message) error

// Messenger will log to STDOUT as well as call a custom log handlers
// defined by the user. If Timestamp is true, then messages are
// prepended with an RFC3339 timestamp.
type Messenger struct {
	closeHandlers []CloseHandler
	handlerMutex  sync.RWMutex
	logHandlers   []MsgHandler
	preprocessor  Preprocessor
	Stdout        bool
	Timestamp     bool
}

// Preprocessor is a function pointer. The Preprocessor is called
// before the message is logged and allows for reformatting of
// messages such as JSON. Set the Discard field to true to drop
// messages.
type Preprocessor func(msg *Message)

// NewFileMessenger will return a new Messenger instance for logging
// to a file. The log file will always show the timestamp, but STDOUT
// will only show the timestamp if Timestamp is true.
func NewFileMessenger(fn string, ts ...bool) (*Messenger, error) {
	var e error
	var file *os.File
	var m *Messenger = NewMessenger(ts...)
	var mutex = sync.Mutex{}

	if len(fn) == 0 {
		return nil, hl.Errorf("No filename provided")
	}

	if file, e = os.Create(fn); e != nil {
		return nil, e
	}

	m.SetCloseHandler(
		func() error {
			var e error

			mutex.Lock()

			if file != nil {
				e = file.Close()
				file = nil
			}

			mutex.Unlock()

			return e
		},
	)

	m.SetMsgHandler(
		func(msg *Message) error {
			var e error

			mutex.Lock()

			if file != nil {
				_, e = file.WriteString(msg.RawString() + "\n")
			}

			mutex.Unlock()

			return e
		},
	)

	return m, nil
}

// NewMessenger will return a new Messenger instance for logging.
func NewMessenger(ts ...bool) *Messenger {
	return &Messenger{
		handlerMutex: sync.RWMutex{},
		Stdout:       true,
		Timestamp:    ((len(ts) > 0) && ts[0]),
	}
}

// AddCloseHandler will add a handler for custom actions when the
// Messenger instance is closed.
func (m *Messenger) AddCloseHandler(handler CloseHandler) {
	m.handlerMutex.Lock()
	m.closeHandlers = append(m.closeHandlers, handler)
	m.handlerMutex.Unlock()
}

// AddMsgHandler will add a handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) AddMsgHandler(handler MsgHandler) {
	m.handlerMutex.Lock()
	m.logHandlers = append(m.logHandlers, handler)
	m.handlerMutex.Unlock()
}

// Close will call the close handler.
func (m *Messenger) Close() error {
	var e error

	for _, f := range m.closeHandlers {
		if e = f(); e != nil {
			return e
		}
	}

	return nil
}

// Debug will log a debug message.
func (m *Messenger) Debug(msg string) error {
	return m.doLog(NewMessage(TypeDebug, msg))
}

// Debugf will log a debug message using a format string.
func (m *Messenger) Debugf(format string, args ...interface{}) error {
	return m.Debug(hl.Sprintf(format, args...))
}

func (m *Messenger) doLog(msg *Message) error {
	var e error

	if m.preprocessor != nil {
		m.handlerMutex.RLock()
		m.preprocessor(msg)
		m.handlerMutex.RUnlock()

		if msg.Discard {
			return nil
		}

		msg.build()
	}

	if m.Stdout {
		if m.Timestamp {
			hl.Println(msg.String())
		} else {
			hl.Println(msg.Text())
		}
	}

	m.handlerMutex.RLock()
	for _, f := range m.logHandlers {
		if e = f(msg); e != nil {
			return e
		}
	}
	m.handlerMutex.RUnlock()

	return nil
}

// Err will log an error message.
func (m *Messenger) Err(msg string) error {
	return m.doLog(NewMessage(TypeErr, msg))
}

// Errf will log an error message using a format string.
func (m *Messenger) Errf(format string, args ...interface{}) error {
	return m.Err(hl.Sprintf(format, args...))
}

// ErrX will log an error message and exit.
func (m *Messenger) ErrX(code int, msg string) {
	m.doLog(NewMessage(TypeErrX, msg))
	os.Exit(code)
}

// ErrfX will log an error message using a format string and exit.
func (m *Messenger) ErrfX(
	code int,
	format string,
	args ...interface{},
) {
	m.ErrX(code, hl.Sprintf(format, args...))
}

// Good will log a good message.
func (m *Messenger) Good(msg string) error {
	return m.doLog(NewMessage(TypeGood, msg))
}

// Goodf will log a good message using a format string.
func (m *Messenger) Goodf(format string, args ...interface{}) error {
	return m.Good(hl.Sprintf(format, args...))
}

// Info will log an info message.
func (m *Messenger) Info(msg string) error {
	return m.doLog(NewMessage(TypeInfo, msg))
}

// Infof will log an info message using a format string.
func (m *Messenger) Infof(format string, args ...interface{}) error {
	return m.Info(hl.Sprintf(format, args...))
}

// Msg will log a message as is.
func (m *Messenger) Msg(msg string) error {
	return m.doLog(NewMessage(TypeMsg, msg))
}

// Msgf will log a message as is using a format string.
func (m *Messenger) Msgf(format string, args ...interface{}) error {
	return m.Msg(hl.Sprintf(format, args...))
}

// SetCloseHandler will set the handler for custom actions when the
// Messenger instance is closed.
func (m *Messenger) SetCloseHandler(handler CloseHandler) {
	m.handlerMutex.Lock()
	m.closeHandlers = []CloseHandler{handler}
	m.handlerMutex.Unlock()
}

// SetColor will disable/enable colors for STDOUT.
func (m *Messenger) SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// SetMsgHandler will set the handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) SetMsgHandler(handler MsgHandler) {
	m.handlerMutex.Lock()
	m.logHandlers = []MsgHandler{handler}
	m.handlerMutex.Unlock()
}

// SetPreprocessor will set the handler for preprocessing messages
// when the Messenger logs a message.
func (m *Messenger) SetPreprocessor(handler Preprocessor) {
	m.handlerMutex.Lock()
	m.preprocessor = handler
	m.handlerMutex.Unlock()
}

// SubInfo will log a subinfo message.
func (m *Messenger) SubInfo(msg string) error {
	return m.doLog(NewMessage(TypeSubInfo, msg))
}

// SubInfof will log a subinfo message using a format string.
func (m *Messenger) SubInfof(
	format string,
	args ...interface{},
) error {
	return m.SubInfo(hl.Sprintf(format, args...))
}

// Warn will log a warn message.
func (m *Messenger) Warn(msg string) error {
	return m.doLog(NewMessage(TypeWarn, msg))
}

// Warnf will log a warn message using a format string.
func (m *Messenger) Warnf(format string, args ...interface{}) error {
	return m.Warn(hl.Sprintf(format, args...))
}
