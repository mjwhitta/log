package log

import (
	"fmt"
	"os"

	hl "github.com/mjwhitta/hilighter"
)

// Debug will log a debug message.
func Debug(msg string) {
	Log(NewMessage(TypeDebug, msg))
}

// Debugf will log a debug message using a format string.
func Debugf(format string, args ...any) {
	Debug(fmt.Sprintf(format, args...))
}

// Err will log an error message.
func Err(msg string) {
	Log(NewMessage(TypeErr, msg))
}

// Errf will log an error message using a format string.
func Errf(format string, args ...any) {
	Err(fmt.Sprintf(format, args...))
}

// ErrX will log an error message and exit.
func ErrX(code int, msg string) {
	Log(NewMessage(TypeErrX, msg))
	os.Exit(code)
}

// ErrXf will log an error message using a format string and exit.
func ErrXf(code int, format string, args ...any) {
	ErrX(code, fmt.Sprintf(format, args...))
}

// Good will log a good message.
func Good(msg string) {
	Log(NewMessage(TypeGood, msg))
}

// Goodf will log a good message using a format string.
func Goodf(format string, args ...any) {
	Good(fmt.Sprintf(format, args...))
}

// Info will log an info message.
func Info(msg string) {
	Log(NewMessage(TypeInfo, msg))
}

// Infof will log an info message using a format string.
func Infof(format string, args ...any) {
	Info(fmt.Sprintf(format, args...))
}

// Log allows for logging of custom message types. If you aren't using
// a custom Messenger, you probably just want to use log.Msg(...)
// instead.
func Log(msg *Message) {
	if Timestamp {
		fmt.Println(msg.String())
	} else {
		fmt.Println(msg.Text())
	}
}

// Msg will log a message as is.
func Msg(msg string) {
	Log(NewMessage(TypeMsg, msg))
}

// Msgf will log a message as is using a format string.
func Msgf(format string, args ...any) {
	Msg(fmt.Sprintf(format, args...))
}

// SetColor will disable/enable colors for stdout.
func SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// SubInfo will log a subinfo message.
func SubInfo(msg string) {
	Log(NewMessage(TypeSubInfo, msg))
}

// SubInfof will log a subinfo message using a format string.
func SubInfof(format string, args ...any) {
	SubInfo(fmt.Sprintf(format, args...))
}

// Warn will log a warn message.
func Warn(msg string) {
	Log(NewMessage(TypeWarn, msg))
}

// Warnf will log a warn message using a format string.
func Warnf(format string, args ...any) {
	Warn(fmt.Sprintf(format, args...))
}
