package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/mjwhitta/errors"
	hl "github.com/mjwhitta/hilighter"
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
// messages such as JSON. Set the Message's Discard field to true to
// drop messages.
type Preprocessor func(msg *Message)

// NewFileMessenger will return a new Messenger instance for logging
// to a file. The log file will always show the timestamp, but STDOUT
// will only show the timestamp if Timestamp is true.
func NewFileMessenger(fn string, ts ...bool) (*Messenger, error) {
	var e error
	var f *os.File
	var m *Messenger = NewMessenger(ts...)
	var mutex sync.Mutex

	if fn == "" {
		return nil, errors.New("no filename provided")
	}

	if f, e = os.Create(filepath.Clean(fn)); e != nil {
		return nil, errors.Newf("failed to create %s: %w", fn, e)
	}

	m.SetCloseHandler(
		func() error {
			var e error

			mutex.Lock()
			defer mutex.Unlock()

			if f != nil {
				e = f.Close()
				f = nil
			}

			if e != nil {
				return errors.Newf("failed to close file: %w", e)
			}

			return nil
		},
	)

	m.SetMsgHandler(
		func(msg *Message) error {
			var e error

			mutex.Lock()
			defer mutex.Unlock()

			if f != nil {
				_, e = f.WriteString(msg.RawString() + "\n")
			}

			if e != nil {
				return errors.Newf("failed to write msg: %w", e)
			}

			return nil
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
	defer m.handlerMutex.Unlock()

	m.closeHandlers = append(m.closeHandlers, handler)
}

// AddMsgHandler will add a handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) AddMsgHandler(handler MsgHandler) {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	m.logHandlers = append(m.logHandlers, handler)
}

// Close will call the close handler.
func (m *Messenger) Close() error {
	for _, f := range m.closeHandlers {
		if e := f(); e != nil {
			return errors.Newf("CloseHandler returned error: %w", e)
		}
	}

	return nil
}

// Debug will log a debug message.
func (m *Messenger) Debug(msg string) error {
	return m.Log(NewMessage(TypeDebug, msg))
}

// Debugf will log a debug message using a format string.
func (m *Messenger) Debugf(format string, args ...any) error {
	return m.Debug(fmt.Sprintf(format, args...))
}

// Err will log an error message.
func (m *Messenger) Err(msg string) error {
	return m.Log(NewMessage(TypeErr, msg))
}

// Errf will log an error message using a format string.
func (m *Messenger) Errf(format string, args ...any) error {
	return m.Err(fmt.Sprintf(format, args...))
}

// ErrX will log an error message and exit.
func (m *Messenger) ErrX(code int, msg string) {
	if e := m.Log(NewMessage(TypeErrX, msg)); e != nil {
		ErrX(code, e.Error())
	}

	os.Exit(code)
}

// ErrXf will log an error message using a format string and exit.
func (m *Messenger) ErrXf(code int, format string, args ...any) {
	m.ErrX(code, fmt.Sprintf(format, args...))
}

// Good will log a good message.
func (m *Messenger) Good(msg string) error {
	return m.Log(NewMessage(TypeGood, msg))
}

// Goodf will log a good message using a format string.
func (m *Messenger) Goodf(format string, args ...any) error {
	return m.Good(fmt.Sprintf(format, args...))
}

// Info will log an info message.
func (m *Messenger) Info(msg string) error {
	return m.Log(NewMessage(TypeInfo, msg))
}

// Infof will log an info message using a format string.
func (m *Messenger) Infof(format string, args ...any) error {
	return m.Info(fmt.Sprintf(format, args...))
}

// Log allows for logging of custom message types.
func (m *Messenger) Log(msg *Message) error {
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
			fmt.Println(msg.String())
		} else {
			fmt.Println(msg.Text())
		}
	}

	m.handlerMutex.RLock()
	defer m.handlerMutex.RUnlock()

	for _, f := range m.logHandlers {
		if e = f(msg); e != nil {
			return errors.Newf("MsgHandler returned error: %w", e)
		}
	}

	return nil
}

// Msg will log a message as is.
func (m *Messenger) Msg(msg string) error {
	return m.Log(NewMessage(TypeMsg, msg))
}

// Msgf will log a message as is using a format string.
func (m *Messenger) Msgf(format string, args ...any) error {
	return m.Msg(fmt.Sprintf(format, args...))
}

// SetCloseHandler will set the handler for custom actions when the
// Messenger instance is closed.
func (m *Messenger) SetCloseHandler(handler CloseHandler) {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	m.closeHandlers = []CloseHandler{handler}
}

// SetColor will disable/enable colors for STDOUT.
func (m *Messenger) SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// SetMsgHandler will set the handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) SetMsgHandler(handler MsgHandler) {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	m.logHandlers = []MsgHandler{handler}
}

// SetPreprocessor will set the handler for preprocessing messages
// when the Messenger logs a message.
func (m *Messenger) SetPreprocessor(handler Preprocessor) {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	m.preprocessor = handler
}

// SubInfo will log a subinfo message.
func (m *Messenger) SubInfo(msg string) error {
	return m.Log(NewMessage(TypeSubInfo, msg))
}

// SubInfof will log a subinfo message using a format string.
func (m *Messenger) SubInfof(format string, args ...any) error {
	return m.SubInfo(fmt.Sprintf(format, args...))
}

// Warn will log a warn message.
func (m *Messenger) Warn(msg string) error {
	return m.Log(NewMessage(TypeWarn, msg))
}

// Warnf will log a warn message using a format string.
func (m *Messenger) Warnf(format string, args ...any) error {
	return m.Warn(fmt.Sprintf(format, args...))
}
