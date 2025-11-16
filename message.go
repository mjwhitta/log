package log

import (
	"maps"
	"time"

	hl "github.com/mjwhitta/hilighter"
)

// Message is struct containing all message related data.
type Message struct {
	Discard bool
	Raw     string
	Type    uint64

	preprocessed string
	text         string
	timestamp    string
}

// NewMessage will return a new Message instance.
func NewMessage(msgType uint64, msg string) (m *Message) {
	var ts string = time.Now().Format(time.RFC3339)

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
	var prefix string

	if Prefixes == nil {
		Prefixes = maps.Clone(defaultPrefixes)
	}

	prefix = Prefixes[m.Type]

	if prefix != "" {
		switch m.Type {
		case TypeDebug:
			prefix = hl.Magenta(prefix)
		case TypeErr, TypeErrX:
			prefix = hl.Red(prefix)
		case TypeGood:
			prefix = hl.Green(prefix)
		case TypeInfo:
			prefix = hl.Blue(prefix)
		case TypeSubInfo:
			prefix = hl.Cyan(prefix)
		case TypeWarn:
			prefix = hl.Yellow(prefix)
		}

		prefix += " "
	}

	m.text = prefix + m.Raw
}

// Preprocessed will return the preprocessed message text.
func (m *Message) Preprocessed() string {
	return m.preprocessed
}

// RawString will return a raw string representation of the Message.
func (m *Message) RawString() string {
	return m.timestamp + ": " + m.Raw
}

// String will return a string representation of the Message.
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
