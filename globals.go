package log

// Timestamp is used to determine whether a timestamp is printed to
// stdout with the message.
var Timestamp bool

const (
	// TypeDebug is a debug message.
	TypeDebug = iota

	// TypeErr is an error message.
	TypeErr

	// TypeErrX is an error message that will exit after.
	TypeErrX

	// TypeGood is a success message.
	TypeGood

	// TypeInfo is an informative message.
	TypeInfo

	// TypeMsg is a generic/plain message.
	TypeMsg

	// TypeSubInfo is an additional info message.
	TypeSubInfo

	// TypeWarn is a warning message.
	TypeWarn
)

// Version is the package version.
const Version = "1.4.6"
