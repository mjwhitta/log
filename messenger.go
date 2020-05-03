package log

import (
	"os"
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// Messenger will log to stdout as well as call a custom log handler
// defined by the user. If Timestamp is true, then messages are
// prepended with an RFC3339 timestamp.
type Messenger struct {
	doneHandler func() error
	logHandler  func(msg string, tsMsg string) error
	Timestamp   bool
}

// NewFileMessenger will return a new Messenger instance for logging
// to a file. The log file will always show the timestamp, but stdout
// will only show the timestamp if Timestamp is true.
func NewFileMessenger(fn string, ts ...bool) (*Messenger, error) {
	var e error
	var file *os.File
	var m *Messenger = NewMessenger(ts...)

	if len(fn) == 0 {
		return nil, hl.Errorf("No filename provided")
	}

	if file, e = os.Create(fn); e != nil {
		return nil, e
	}

	m.SetDoneHandler(
		func() error {
			if file != nil {
				return file.Close()
			}

			return nil
		},
	)

	m.SetLogHandler(
		func(msg string, tsMsg string) error {
			var e error

			if file != nil {
				_, e = file.WriteString(hl.Plain(tsMsg) + "\n")
				return e
			}

			return nil
		},
	)

	return m, nil
}

// NewMessenger will return a new Messenger instance for logging.
func NewMessenger(ts ...bool) *Messenger {
	return &Messenger{Timestamp: ((len(ts) > 0) && ts[0])}
}

// Close will call the done handler.
func (m *Messenger) Close() error {
	if m.doneHandler != nil {
		return m.doneHandler()
	}

	return nil
}

// SetColor will disable/enable colors for stdout.
func (m *Messenger) SetColor(enabled bool) {
	hl.Disable(!enabled)
}

func (m *Messenger) doLog(msg string) error {
	var ts = time.Now().Format(time.RFC3339) + ": "
	var tsMsg = ts + msg

	if m.Timestamp {
		hl.Println(tsMsg)
	} else {
		hl.Println(msg)
	}

	if m.logHandler != nil {
		return m.logHandler(msg, tsMsg)
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

// SetDoneHandler will set the handler for when the Messenger instance
// is closed.
func (m *Messenger) SetDoneHandler(handler func() error) {
	m.doneHandler = handler
}

// SetLogHandler will set the handler for custom actions when the
// Messenger logs a message.
func (m *Messenger) SetLogHandler(
	handler func(msg string, tsMsg string) error,
) {
	m.logHandler = handler
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
