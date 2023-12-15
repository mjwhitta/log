package log

// Timestamp is used to determine whether a timestamp is printed to
// stdout with the message.
var Timestamp bool

// Consts for log message types
const (
	TypeDebug   = iota // TypeDebug is a debug message
	TypeErr            // TypeErr is an error message
	TypeErrX           // TypeErrX is an error message that will exit
	TypeGood           // TypeGood is a success message
	TypeInfo           // TypeInfo is an informative message
	TypeMsg            // TypeMsg is a generic/plain message
	TypeSubInfo        // TypeSubInfo is an additional info message
	TypeWarn           // TypeWarn is a warning message
)

// Version is the package version.
const Version = "1.6.5"
