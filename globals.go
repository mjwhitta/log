package log

// Timestamp is used to determine whether a timestamp is printed to
// stdout with the message.
var Timestamp bool

// TypeDebug is a debug message.
const TypeDebug uint8 = 0

// TypeErr is an error message.
const TypeErr uint8 = TypeDebug + 1

// TypeErrX is an error message that will exit after.
const TypeErrX uint8 = TypeErr + 1

// TypeGood is a success message.
const TypeGood uint8 = TypeErrX + 1

// TypeInfo is an informative message.
const TypeInfo uint8 = TypeGood + 1

// TypeMsg is a generic/plain message.
const TypeMsg uint8 = TypeInfo + 1

// TypeSubInfo is an additional info message.
const TypeSubInfo uint8 = TypeMsg + 1

// TypeWarn is a warning message.
const TypeWarn uint8 = TypeSubInfo + 1

// Version is the package version.
const Version = "1.3.6"
