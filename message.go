package log

import (
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// Message is struct containing all message related data.
type Message struct {
	Text      string
	Timestamp string
	TimeText  string
	Type      uint8
}

// NewMessage will return a new Message instance.
func NewMessage(msgType uint8, msg string) Message {
	var ts = time.Now().Format(time.RFC3339)

	switch msgType {
	case TypeDebug:
		msg = hl.Magenta("[#] " + msg)
	case TypeErr, TypeErrX:
		msg = hl.Red("[!] " + msg)
	case TypeGood:
		msg = hl.Green("[+] " + msg)
	case TypeInfo:
		msg = hl.White("[*] " + msg)
	case TypeSubInfo:
		msg = hl.Cyan("[=] " + msg)
	case TypeWarn:
		msg = hl.Yellow("[-] " + msg)
	}

	return Message{
		Text:      msg,
		Timestamp: ts,
		TimeText:  ts + ": " + msg,
		Type:      msgType,
	}
}
