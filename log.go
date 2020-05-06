package log

import (
	"os"

	hl "gitlab.com/mjwhitta/hilighter"
)

// Debug will log a debug message.
func Debug(msg string) {
	doLog(NewMessage(TypeDebug, msg))
}

// Debugf will log a debug message using a format string.
func Debugf(format string, args ...interface{}) {
	Debug(hl.Sprintf(format, args...))
}

func doLog(msg Message) {
	if Timestamp {
		hl.Println(msg.TimeText)
	} else {
		hl.Println(msg.Text)
	}
}

// Err will log an error message.
func Err(msg string) {
	doLog(NewMessage(TypeErr, msg))
}

// Errf will log an error message using a format string.
func Errf(format string, args ...interface{}) {
	Err(hl.Sprintf(format, args...))
}

// ErrX will log an error message and exit.
func ErrX(code int, msg string) {
	doLog(NewMessage(TypeErrX, msg))
	os.Exit(code)
}

// ErrfX will log an error message using a format string and exit.
func ErrfX(code int, format string, args ...interface{}) {
	ErrX(code, hl.Sprintf(format, args...))
}

// Good will log a good message.
func Good(msg string) {
	doLog(NewMessage(TypeGood, msg))
}

// Goodf will log a good message using a format string.
func Goodf(format string, args ...interface{}) {
	Good(hl.Sprintf(format, args...))
}

// Info will log an info message.
func Info(msg string) {
	doLog(NewMessage(TypeInfo, msg))
}

// Infof will log an info message using a format string.
func Infof(format string, args ...interface{}) {
	Info(hl.Sprintf(format, args...))
}

// Msg will log a message as is.
func Msg(msg string) {
	doLog(NewMessage(TypeMsg, msg))
}

// Msgf will log a message as is using a format string.
func Msgf(format string, args ...interface{}) {
	Msg(hl.Sprintf(format, args...))
}

// SetColor will disable/enable colors for stdout.
func SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// SubInfo will log a subinfo message.
func SubInfo(msg string) {
	doLog(NewMessage(TypeSubInfo, msg))
}

// SubInfof will log a subinfo message using a format string.
func SubInfof(format string, args ...interface{}) {
	SubInfo(hl.Sprintf(format, args...))
}

// Warn will log a warn message.
func Warn(msg string) {
	doLog(NewMessage(TypeWarn, msg))
}

// Warnf will log a warn message using a format string.
func Warnf(format string, args ...interface{}) {
	Warn(hl.Sprintf(format, args...))
}
