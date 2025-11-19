package log

import "maps"

// Version is the package version.
const Version string = "1.8.3"

// Consts for log message types
//
//nolint:grouper // This is an iota block
const (
	TypeDebug uint64 = iota
	TypeErr
	TypeErrX // An error message that will also exit
	TypeGood
	TypeInfo
	TypeMsg // Generic message
	TypeSubInfo
	TypeWarn
)

var (
	// Prefixes allows you to customize each log message prefix.
	Prefixes map[uint64]string = map[uint64]string{
		TypeDebug:   "[#]",
		TypeErr:     "[!]",
		TypeErrX:    "[!]",
		TypeGood:    "[+]",
		TypeInfo:    "[*]",
		TypeSubInfo: "[=]",
		TypeWarn:    "[-]",
	}

	// Timestamp is used to determine whether a timestamp is printed
	// to stdout with the message.
	Timestamp bool

	defaultPrefixes map[uint64]string
)

func init() {
	defaultPrefixes = maps.Clone(Prefixes)
}
