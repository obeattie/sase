package query

import (
	"bytes"
	"fmt"
	"time"

	"github.com/obeattie/sase/domain"
)

type EventCapture interface {
	Representable
	// Names returns a mapping of (local alias: type) for all captured events
	Names() map[string]string
	// Matches determines whether the passed event should be captured, returning local name(s) it should be capture as.
	Matches(e domain.Event) []string
	// Negations returns the local aliases of all negated events
	Negations() []string
	// aliases returns just the event aliases used (including duplicates)
	aliases() []string
	// evaluate allows the capture to participate in the matching phase, primarily so it may veto a match
	evaluate(domain.CapturedEvents) Result
}

// A basicEventCapture represents a capture of a single event, determined by its type
type basicEventCapture struct {
	eventType string
	name      string
}

func (c *basicEventCapture) Matches(e domain.Event) []string {
	if c.eventType == e.Type() {
		return []string{c.name}
	} else {
		return nil
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

func (c *basicEventCapture) evaluate(evs domain.CapturedEvents) Result {
	if e, ok := evs[c.name]; !ok { // Not yet matched
		return Uncertain
	} else if e.Type() != c.eventType { // Incorrect type
		return Negative
	} else { // Correct type
		return Positive
	}
}

// A seqEventCapture captures a sequence of events
type seqEventCapture []EventCapture

func (c seqEventCapture) Matches(e domain.Event) []string {
	var (
		result         []string
		negatedResult  []string
		negations      = c.Negations()
		negatedAliases = make(map[string]struct{}, len(negations))
	)
	for _, negatedAlias := range negations {
		negatedAliases[negatedAlias] = struct{}{}
	}
	for _, subCap := range c {
		for _, alias := range subCap.Matches(e) {
			if _, ok := negatedAliases[alias]; ok {
				negatedResult = append(negatedResult, alias)
			} else {
				result = append(result, alias)
			}
		}
	}
	// Negated aliases must be at the end (this is important, as negations must be evaluated later than positive
	// matches, or we may inadvertently preempt a valid stack)
	return append(result, negatedResult...)
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
	buf.WriteRune(')')
	return buf.String()
}

func (c seqEventCapture) Negations() []string {
	var result []string
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

func (c seqEventCapture) evaluate(evs domain.CapturedEvents) Result {
	result := Uncertain
	for i, subCap := range c {
		if sr := subCap.evaluate(evs); i == 0 {
			result = sr
		} else {
			result = result.And(sr)
		}
	}

	if result != Negative { // Ensure the captures are actually in sequence
		var lastTs time.Time
		for _, alias := range c.aliases() {
			if ev, ok := evs[alias]; ok {
				when := ev.When()
				if when.Before(lastTs) {
					return Negative
				}
				lastTs = when
			}
		}
	}

	return result
}

// An anyEventCapture will capture a match of any of its contained captures
type anyEventCapture []EventCapture

func (c anyEventCapture) Matches(e domain.Event) []string {
	var (
		result         []string
		negatedResult  []string
		negations      = c.Negations()
		negatedAliases = make(map[string]struct{}, len(negations))
	)
	for _, negatedAlias := range negations {
		negatedAliases[negatedAlias] = struct{}{}
	}
	for _, subCap := range c {
		for _, alias := range subCap.Matches(e) {
			if _, ok := negatedAliases[alias]; ok {
				negatedResult = append(negatedResult, alias)
			} else {
				result = append(result, alias)
			}
		}
	}
	// Negated aliases must be at the end (this is important, as negations must be evaluated later than positive
	// matches, or we may inadvertently preempt a valid stack)
	return append(result, negatedResult...)
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
	var result []string
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

func (c anyEventCapture) evaluate(evs domain.CapturedEvents) Result {
	result := Uncertain
	for i, subCap := range c {
		if sr := subCap.evaluate(evs); i == 0 {
			result = sr
		} else {
			result = result.Or(sr)
		}
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

func (c *negatedEventCapture) evaluate(evs domain.CapturedEvents) Result {
	r := c.EventCapture.evaluate(evs)
	switch r {
	case Positive:
		return Invalid
	case Uncertain:
		return Positive
	default:
		return r
	}
}
