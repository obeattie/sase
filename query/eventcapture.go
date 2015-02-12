package query

import (
	"bytes"
	"fmt"

	"github.com/obeattie/sase/domain"
)

type EventCapture interface {
	Representable
	// Names returns a mapping of (local aliase: type) for all captured events
	Names() map[string]string
	// Matches determines whether the passed event should be captured. If so, the local name it should be captured
	// under is returned. If not, an empty string is returned.
	Matches(e domain.Event) string
	// Negations returns the local aliases of all negated events
	Negations() []string
	// aliases returns just the event aliases used (including duplicates)
	aliases() []string
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

func (c *basicEventCapture) Negations() []string {
	return nil
}

func (c *basicEventCapture) Names() map[string]string {
	return map[string]string{
		c.name: c.eventType,
	}
}

func (c *basicEventCapture) aliases() []string {
	return []string{c.name}
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

func (c seqEventCapture) Negations() []string {
	result := make([]string, 0)
	for _, subCap := range c {
		result = append(result, subCap.Negations()...)
	}
	return result
}

func (c seqEventCapture) Names() map[string]string {
	result := make(map[string]string, len(c))
	for _, subCap := range c {
		for k, v := range subCap.Names() {
			result[k] = v
		}
	}
	return result
}

func (c seqEventCapture) aliases() []string {
	result := make([]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.aliases()...)
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

func (c anyEventCapture) Negations() []string {
	result := make([]string, 0)
	for _, subCap := range c {
		result = append(result, subCap.Negations()...)
	}
	return result
}

func (c anyEventCapture) Names() map[string]string {
	result := make(map[string]string, len(c))
	for _, subCap := range c {
		for k, v := range subCap.Names() {
			result[k] = v
		}
	}
	return result
}

func (c anyEventCapture) aliases() []string {
	result := make([]string, 0, len(c))
	for _, subCap := range c {
		result = append(result, subCap.aliases()...)
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

func (c *negatedEventCapture) Negations() []string {
	subNegations := c.EventCapture.Negations()
	result := make([]string, 0)
candidateLoop:
	for alias, _ := range c.Names() {
		for _, negatedAlias := range subNegations {
			if alias == negatedAlias {
				continue candidateLoop
			}
		}
		result = append(result, alias)
	}
	return result
}
