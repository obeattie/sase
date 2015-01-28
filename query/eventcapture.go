package query

import (
	"bytes"
	"fmt"

	"github.com/obeattie/sase/domain"
)

type EventCapture interface {
	Representable
	// Names returns the types and local aliases of all captured events
	Names() [][2]string
	// Matches determines whether the passed event should be captured. If so, the local name it should be captured
	// under is returned. If not, an empty string is returned.
	Matches(e domain.Event) string
	// Negations returns the types and local aliases of captures that are negated
	Negations() [][2]string
}

// A basicEventCapture represents a capture of a single event, determined by its type
type basicEventCapture struct {
	eventType string
	name      string
}

func (c *basicEventCapture) Matches(e domain.Event) string {
	if c.eventType == e.Type() {
		return c.name
	} else {
		return ""
	}
}

func (c *basicEventCapture) QueryText() string {
	return fmt.Sprintf("%s %s", c.eventType, c.name)
}

func (c *basicEventCapture) Negations() [][2]string {
	return nil
}

func (c *basicEventCapture) Names() [][2]string {
	return [][2]string{{c.eventType, c.name}}
}

// A seqEventCapture captures a sequence of events
type seqEventCapture []EventCapture

func (c seqEventCapture) Matches(e domain.Event) string {
	for _, subCap := range c {
		if alias := subCap.Matches(e); alias != "" {
			return alias
		}
	}
	return ""
}

func (c seqEventCapture) QueryText() string {
	buf := new(bytes.Buffer)
	buf.WriteString("SEQ(")
	for i, subCap := range c {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(subCap.QueryText())
	}
	buf.WriteString(")")
	return buf.String()
}

func (c seqEventCapture) Negations() [][2]string {
	result := make([][2]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.Negations()...)
	}
	return result
}

func (c seqEventCapture) Names() [][2]string {
	result := make([][2]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.Names()...)
	}
	return result
}

// An anyEventCapture will capture a match of any of its contained captures
type anyEventCapture []EventCapture

func (c anyEventCapture) Matches(e domain.Event) string {
	for _, subCap := range c {
		if alias := subCap.Matches(e); alias != "" {
			return alias
		}
	}
	return ""
}

func (c anyEventCapture) QueryText() string {
	buf := new(bytes.Buffer)
	buf.WriteString("ANY(")
	for i, subCap := range c {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(subCap.QueryText())
	}
	buf.WriteString(")")
	return buf.String()
}

func (c anyEventCapture) Negations() [][2]string {
	result := make([][2]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.Negations()...)
	}
	return result
}

func (c anyEventCapture) Names() [][2]string {
	result := make([][2]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.Names()...)
	}
	return result
}

// A negatedEventCapture still captures the event, but reports it as being a negated event (this is significant in later
// processing)
type negatedEventCapture struct {
	EventCapture
}

func (c *negatedEventCapture) QueryText() string {
	return fmt.Sprintf("!(%s)", c.EventCapture.QueryText())
}

func (c *negatedEventCapture) Negations() [][2]string {
	subNegations := c.EventCapture.Negations()
	result := make([][2]string, 0)

candidateLoop:
	for _, nameSpec := range c.Names() {
		for _, negSpec := range subNegations {
			if nameSpec[0] == negSpec[0] && nameSpec[1] == negSpec[1] {
				continue candidateLoop
			}
			result = append(result, nameSpec)
		}
	}

	return result
}
