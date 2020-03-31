package log

import hl "gitlab.com/mjwhitta/hilighter"

// Err will log an error message.
func Err(msg string) {
	logToFile(hl.Redf("[!] %s", msg), Timestamp)
}

// Errf will log an error message using a format string.
func Errf(fmt string, args ...interface{}) {
	Err(hl.Sprintf(fmt, args...))
}

// Good will log a good message.
func Good(msg string) {
	logToFile(hl.Greenf("[+] %s", msg), Timestamp)
}

// Goodf will log a good message using a format string.
func Goodf(fmt string, args ...interface{}) {
	Good(hl.Sprintf(fmt, args...))
}

// Info will log an info message.
func Info(msg string) {
	logToFile(hl.Whitef("[*] %s", msg), Timestamp)
}

// Infof will log an info message using a format string.
func Infof(fmt string, args ...interface{}) {
	Info(hl.Sprintf(fmt, args...))
}

// Msg will log a message as is.
func Msg(msg string) {
	logToFile(msg, Timestamp)
}

// Msgf will log a message as is using a format string.
func Msgf(fmt string, args ...interface{}) {
	Msg(hl.Sprintf(fmt, args...))
}

// SubInfo will log a subinfo message.
func SubInfo(msg string) {
	logToFile(hl.Cyanf("[=] %s", msg), Timestamp)
}

// SubInfof will log a subinfo message using a format string.
func SubInfof(fmt string, args ...interface{}) {
	SubInfo(hl.Sprintf(fmt, args...))
}

// Warn will log a warn message.
func Warn(msg string) {
	logToFile(hl.Yellowf("[-] %s", msg), Timestamp)
}

// Warnf will log a warn message using a format string.
func Warnf(fmt string, args ...interface{}) {
	Warn(hl.Sprintf(fmt, args...))
}
