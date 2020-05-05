package log

// Timestamp is used to determine whether a timestamp is printed to
// stdout with the message.
var Timestamp bool

// TypeErr is an error message.
const TypeErr uint8 = 0

// TypeErrX is an error message that will exit after.
const TypeErrX uint8 = 1

// TypeGood is a success message.
const TypeGood uint8 = 2

// TypeInfo is an informative message.
const TypeInfo uint8 = 3

// TypeMsg is a generic/plain message.
const TypeMsg uint8 = 4

// TypeSubInfo is an additional info message.
const TypeSubInfo uint8 = 5

// TypeWarn is a warning message.
const TypeWarn uint8 = 6

// Version is the package version.
const Version = "1.3.2"
