package query

import (
	"bytes"
	"time"
)

// A Representable may be converted back to a string query
type Representable interface {
	Query() string
}

type Query struct {
	// Capture represents the types of events that should be captured
	Capture EventCapture
	// Filter represents the filters that should be applied to the captured events (referring to the events by their
	// local aliases)
	Filter *bool // TODO: Implement
	// Window represents the interval over which events may be matched
	Window time.Duration
}

func (q *Query) Query() string {
	buf := new(bytes.Buffer)
	buf.WriteString("EVENT")
	if q.Capture != nil {
		buf.WriteRune(' ')
		buf.WriteString(q.Capture.Query())
	}
	if q.Filter != nil {
		buf.WriteString(" WHERE ")
		// TODO: Implement
	}
	if q.Window != 0 {
		buf.WriteString(" WITHIN ")
		buf.WriteString(q.Window.String())
	}
	buf.WriteString(";")
	return buf.String()
}
