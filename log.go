package log

import (
	hl "gitlab.com/mjwhitta/hilighter"
)

func doLog(msg Message) {
	if Timestamp {
		hl.Println(msg.TimeText)
	} else {
		hl.Println(msg.Text)
	}
}

// SetColor will disable/enable colors for stdout.
func SetColor(enabled bool) {
	hl.Disable(!enabled)
}

// Err will log an error message.
func Err(msg string) {
	doLog(NewMessage(TypeErr, msg))
}

// Errf will log an error message using a format string.
func Errf(format string, args ...interface{}) {
	Err(hl.Sprintf(format, args...))
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
