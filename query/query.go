package query

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/obeattie/sase/domain"
)

// CapturedEvents represents an event sequence that has been captured from an input stream (or streams) by a query.
type CapturedEvents map[string]domain.Event

// A Representable may be converted back to a string query
type Representable interface {
	QueryText() string
}

type Query struct {
	// capture represents the types of events that should be captured
	capture EventCapture
	// predicate represents filter(s) to be applied to a captured event stream
	predicate Predicate
	// window represents the interval over which events may be matched
	window time.Duration
}

func (q *Query) QueryText() string {
	buf := new(bytes.Buffer)
	buf.WriteString("EVENT")
	if q.capture != nil {
		buf.WriteRune(' ')
		buf.WriteString(q.capture.QueryText())
	}
	if q.predicate != nil {
		buf.WriteString(" WHERE ")
		buf.WriteString(q.predicate.QueryText())
	}
	if q.window != 0 {
		buf.WriteString(" WITHIN ")
		buf.WriteString(q.window.String())
	}
	buf.WriteString(";")
	return buf.String()
}

func (q *Query) Window() time.Duration {
	return q.window
}

// Returns whether this query is interested in the passed event
func (q *Query) ShouldCapture(e domain.Event) bool {
	return q.capture.Matches(e) != ""
}

func (q *Query) validate() error {
	if len(q.capture.Names()) < 1 {
		return fmt.Errorf("Query must have at least one event capture")
	}

	// Check for overlapping aliases
	seenAliases := make(map[string]bool)
	duplicateAliases := make([]string, 0)
	for _, captureName := range q.capture.Names() {
		alias := captureName[1]
		if _, ok := seenAliases[alias]; ok {
			duplicateAliases = append(duplicateAliases, alias)
		} else {
			seenAliases[alias] = true
		}
	}
	if len(duplicateAliases) > 0 {
		return fmt.Errorf("Query has duplicate aliases %s", strings.Join(duplicateAliases, ", "))
	}

	if q.window < 0 {
		return fmt.Errorf("Query has negative window duration")
	}

	return nil
}
