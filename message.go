package log

import (
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// Message is struct containing all message related data.
type Message struct {
	Raw       string
	Text      string
	Timestamp string
	TimeText  string
	Type      uint8
}

// NewMessage will return a new Message instance.
func NewMessage(msgType uint8, msg string) Message {
	var formatted string
	var ts = time.Now().Format(time.RFC3339)

	switch msgType {
	case TypeDebug:
		formatted = hl.Magenta("[#] " + msg)
	case TypeErr, TypeErrX:
		formatted = hl.Red("[!] " + msg)
	case TypeGood:
		formatted = hl.Green("[+] " + msg)
	case TypeInfo:
		formatted = hl.White("[*] " + msg)
	case TypeSubInfo:
		formatted = hl.Cyan("[=] " + msg)
	case TypeWarn:
		formatted = hl.Yellow("[-] " + msg)
	default:
		formatted = msg
	}

	return Message{
		Raw:       msg,
		Text:      formatted,
		Timestamp: ts,
		TimeText:  ts + ": " + formatted,
		Type:      msgType,
	}
}

func (m Message) String() string {
	return m.TimeText
}
