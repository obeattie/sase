package query

import (
	"bytes"
	"fmt"

	"github.com/obeattie/sase-parser/domain"
)

type EventCapture interface {
	Representable
	// Matches determines whether the passed event should be captured
	Matches(e domain.Event) bool
}

type multiEventCapture interface {
	add(EventCapture)
}

// A basicEventCapture represents a capture of a single event, determined by its type
type basicEventCapture struct {
	eventType string
	name      string
}

func (c *basicEventCapture) Matches(e domain.Event) bool {
	return c.eventType == e.Type()
}

func (c *basicEventCapture) Query() string {
	return fmt.Sprintf("%s %s", c.eventType, c.name)
}

// A seqEventCapture captures a sequence of events
type seqEventCapture []EventCapture

func (c seqEventCapture) Matches(e domain.Event) bool {
	for _, subCap := range c {
		if subCap.Matches(e) {
			return true
		}
	}
	return false
}

func (c seqEventCapture) Query() string {
	buf := new(bytes.Buffer)
	buf.WriteString("SEQ(")
	for i, subCap := range c {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(subCap.Query())
	}
	buf.WriteString(")")
	return buf.String()
}

// An anyEventCapture will capture a match of any of its contained captures
type anyEventCapture []EventCapture

func (c anyEventCapture) Query() string {
	buf := new(bytes.Buffer)
	buf.WriteString("ANY(")
	for i, subCap := range c {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(subCap.Query())
	}
	buf.WriteString(")")
	return buf.String()
}

func (c anyEventCapture) Matches(e domain.Event) bool {
	for _, subCap := range c {
		if subCap.Matches(e) {
			return true
		}
	}
	return false
}

// A negatedEventCapture will capture a match of any event other than its contained capture
type negatedEventCapture struct {
	EventCapture
}

func (c *negatedEventCapture) Matches(e domain.Event) bool {
	return !c.EventCapture.Matches(e)
}

func (c *negatedEventCapture) Query() string {
	return fmt.Sprintf("!(%s)", c.EventCapture.Query())
}
