package log

import (
	"time"

	hl "github.com/mjwhitta/hilighter"
)

// Message is struct containing all message related data.
type Message struct {
	Discard      bool
	preprocessed string
	Raw          string
	text         string
	timestamp    string
	Type         uint8
}

// NewMessage will return a new Message instance.
func NewMessage(msgType uint8, msg string) (m *Message) {
	var ts = time.Now().Format(time.RFC3339)

	m = &Message{
		preprocessed: msg,
		Raw:          msg,
		timestamp:    ts,
		Type:         msgType,
	}
	m.build()

	return
}

func (m *Message) build() {
	switch m.Type {
	case TypeDebug:
		m.text = hl.Magenta("[#] " + m.Raw)
	case TypeErr, TypeErrX:
		m.text = hl.Red("[!] " + m.Raw)
	case TypeGood:
		m.text = hl.Green("[+] " + m.Raw)
	case TypeInfo:
		m.text = hl.White("[*] " + m.Raw)
	case TypeSubInfo:
		m.text = hl.Cyan("[=] " + m.Raw)
	case TypeWarn:
		m.text = hl.Yellow("[-] " + m.Raw)
	default:
		m.text = m.Raw
	}
}

// Preprocessed will return the preprocessed message text.
func (m *Message) Preprocessed() string {
	return m.preprocessed
}

// RawString will return the raw string representation of the Message
// instance.
func (m *Message) RawString() string {
	return m.timestamp + ": " + m.Raw
}

// String will return the string representation of the Message
// instance.
func (m *Message) String() string {
	return m.timestamp + ": " + m.text
}

// Text will return the processed message text (w/ no timestamp)
func (m *Message) Text() string {
	return m.text
}

// Timestamp will return the timestamp of the message.
func (m *Message) Timestamp() string {
	return m.timestamp
}
