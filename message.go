package log

import (
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// Message is struct containing all message related data.
type Message struct {
	Preprocessed string
	Raw          string
	Text         string
	Timestamp    string
	Type         uint8
}

// NewMessage will return a new Message instance.
func NewMessage(msgType uint8, msg string) (m Message) {
	var ts = time.Now().Format(time.RFC3339)

	m = Message{
		Preprocessed: msg,
		Timestamp:    ts,
		Type:         msgType,
	}
	m.build(msg)

	return
}

func (m Message) build(raw string) {
	m.Raw = raw

	switch m.Type {
	case TypeDebug:
		m.Text = hl.Magenta("[#] " + m.Raw)
	case TypeErr, TypeErrX:
		m.Text = hl.Red("[!] " + m.Raw)
	case TypeGood:
		m.Text = hl.Green("[+] " + m.Raw)
	case TypeInfo:
		m.Text = hl.White("[*] " + m.Raw)
	case TypeSubInfo:
		m.Text = hl.Cyan("[=] " + m.Raw)
	case TypeWarn:
		m.Text = hl.Yellow("[-] " + m.Raw)
	default:
		m.Text = m.Raw
	}
}

// RawString will return the raw string representation of the Message
// instance.
func (m Message) RawString() string {
	return m.Timestamp + ": " + m.Raw
}

// String will return the string representation of the Message
// instance.
func (m Message) String() string {
	return m.Timestamp + ": " + m.Text
}
