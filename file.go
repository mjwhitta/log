package log

import (
	"os"

	hl "gitlab.com/mjwhitta/hilighter"
)

// File is a struct containing info to log to a file. By default it
// will log to stdout. If a filename is configured, it will also log
// to that file.
type File struct {
	file      *os.File
	Timestamp bool
}

// New will return a new File instance for logging.
func New(filename ...string) (*File, error) {
	var e error
	var file *os.File

	if (len(filename) > 0) && (len(filename[0]) > 0) {
		if file, e = os.Create(filename[0]); e != nil {
			return nil, e
		}
	}

	return &File{file: file}, nil
}

// Close will close the log file.
func (f *File) Close() {
	if f.file != nil {
		f.file.Close()
	}
}

// Err will log an error message.
func (f *File) Err(msg string) {
	logToFile(hl.Redf("[!] %s", msg), f.Timestamp, f.file)
}

// Errf will log an error message using a format string.
func (f *File) Errf(fmt string, args ...interface{}) {
	f.Err(hl.Sprintf(fmt, args...))
}

// Good will log a good message.
func (f *File) Good(msg string) {
	logToFile(hl.Greenf("[+] %s", msg), f.Timestamp, f.file)
}

// Goodf will log a good message using a format string.
func (f *File) Goodf(fmt string, args ...interface{}) {
	f.Good(hl.Sprintf(fmt, args...))
}

// Info will log an info message.
func (f *File) Info(msg string) {
	logToFile(hl.Whitef("[*] %s", msg), f.Timestamp, f.file)
}

// Infof will log an info message using a format string.
func (f *File) Infof(fmt string, args ...interface{}) {
	f.Info(hl.Sprintf(fmt, args...))
}

// Msg will log a message as is.
func (f *File) Msg(msg string) {
	logToFile(msg, f.Timestamp, f.file)
}

// Msgf will log a message as is using a format string.
func (f *File) Msgf(fmt string, args ...interface{}) {
	f.Msg(hl.Sprintf(fmt, args...))
}

// SubInfo will log a subinfo message.
func (f *File) SubInfo(msg string) {
	logToFile(hl.Cyanf("[=] %s", msg), f.Timestamp, f.file)
}

// SubInfof will log a subinfo message using a format string.
func (f *File) SubInfof(fmt string, args ...interface{}) {
	f.SubInfo(hl.Sprintf(fmt, args...))
}

// Warn will log a warn message.
func (f *File) Warn(msg string) {
	logToFile(hl.Yellowf("[-] %s", msg), f.Timestamp, f.file)
}

// Warnf will log a warn message using a format string.
func (f *File) Warnf(fmt string, args ...interface{}) {
	f.Warn(hl.Sprintf(fmt, args...))
}
