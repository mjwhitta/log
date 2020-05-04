package log

import (
	"os"
	"sync"
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// CustomDoneHandler is a function pointer.
type CustomDoneHandler func() error

// CustomLogHandler is a function pointer.
type CustomLogHandler func(ts string, msg string, tsMsg string) error

// Messenger will log to STDOUT as well as call a custom log handlers
// defined by the user. If Timestamp is true, then messages are
// prepended with an RFC3339 timestamp.
type Messenger struct {
	doneHandlers []CustomDoneHandler
	logHandlers  []CustomLogHandler
	Stdout       bool
	Timestamp    bool
}

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

	m.SetDoneHandler(
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

	m.SetLogHandler(
		func(ts string, msg string, tsMsg string) error {
			var e error

			mutex.Lock()

			if file != nil {
				_, e = file.WriteString(hl.Plain(tsMsg) + "\n")
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
		Stdout:    true,
		Timestamp: ((len(ts) > 0) && ts[0]),
	}
}

// AddDoneHandler will add a handler for custom actions when the
// Messenger instance is closed.
func (m *Messenger) AddDoneHandler(handler CustomDoneHandler) {
	m.doneHandlers = append(m.doneHandlers, handler)
}

// AddLogHandler will add a handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) AddLogHandler(handler CustomLogHandler) {
	m.logHandlers = append(m.logHandlers, handler)
}

// Close will call the done handler.
func (m *Messenger) Close() error {
	var e error

	for _, f := range m.doneHandlers {
		if e = f(); e != nil {
			return e
		}
	}

	return nil
}

func (m *Messenger) doLog(msg string) error {
	var e error
	var ts = time.Now().Format(time.RFC3339)
	var tsMsg = ts + ": " + msg

	if m.Stdout {
		if m.Timestamp {
			hl.Println(tsMsg)
		} else {
			hl.Println(msg)
		}
	}

	for _, f := range m.logHandlers {
		if e = f(ts, msg, tsMsg); e != nil {
			return e
		}
	}

	return nil
}

// Err will log an error message.
func (m *Messenger) Err(msg string) error {
	return m.doLog(hl.Red("[!] " + msg))
}

// Errf will log an error message using a format string.
func (m *Messenger) Errf(format string, args ...interface{}) error {
	return m.Err(hl.Sprintf(format, args...))
}

// Good will log a good message.
func (m *Messenger) Good(msg string) error {
	return m.doLog(hl.Green("[+] " + msg))
}

// Goodf will log a good message using a format string.
func (m *Messenger) Goodf(format string, args ...interface{}) error {
	return m.Good(hl.Sprintf(format, args...))
}

// Info will log an info message.
func (m *Messenger) Info(msg string) error {
	return m.doLog(hl.White("[*] " + msg))
}

// Infof will log an info message using a format string.
func (m *Messenger) Infof(format string, args ...interface{}) error {
	return m.Info(hl.Sprintf(format, args...))
}

// Msg will log a message as is.
func (m *Messenger) Msg(msg string) error {
	return m.doLog(msg)
}

// Msgf will log a message as is using a format string.
func (m *Messenger) Msgf(format string, args ...interface{}) error {
	return m.Msg(hl.Sprintf(format, args...))
}

// SetColor will disable/enable colors for STDOUT.
func (m *Messenger) SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// SetDoneHandler will set the handler for custom actions when the
// Messenger instance is closed.
func (m *Messenger) SetDoneHandler(handler CustomDoneHandler) {
	m.doneHandlers = []CustomDoneHandler{handler}
}

// SetLogHandler will set the handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) SetLogHandler(handler CustomLogHandler) {
	m.logHandlers = []CustomLogHandler{handler}
}

// SubInfo will log a subinfo message.
func (m *Messenger) SubInfo(msg string) error {
	return m.doLog(hl.Cyan("[=] " + msg))
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
	return m.doLog(hl.Yellow("[-] " + msg))
}

// Warnf will log a warn message using a format string.
func (m *Messenger) Warnf(format string, args ...interface{}) error {
	return m.Warn(hl.Sprintf(format, args...))
}
