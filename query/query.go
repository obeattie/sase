package query

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/obeattie/sase/domain"
)

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

// ShouldCapture returns whether this query is interested in the passed event. If so, the alias under which it should be
// captured is returned.
func (q *Query) ShouldCapture(e domain.Event) string {
	return q.capture.Matches(e)
}

// Captures returns a mapping of (local alias: type) for all captured events
func (q *Query) Captures() map[string]string {
	return q.capture.Names()
}

func (q *Query) Evaluate(evs domain.CapturedEvents) PredicateResult {
	if q.predicate == nil {
		return PredicateResultPositive
	}
	return q.predicate.Evaluate(evs)
}

func (q *Query) validate() error {
	if len(q.capture.Names()) < 1 {
		return fmt.Errorf("Query must have at least one event capture")
	}

	// Check for overlapping aliases
	seenAliases, duplicateAliases := make(map[string]bool), make([]string, 0)
	for _, alias := range q.capture.aliases() {
		if _, ok := seenAliases[alias]; ok {
			duplicateAliases = append(duplicateAliases, alias)
		}
		seenAliases[alias] = true
	}
	if len(duplicateAliases) > 0 {
		return fmt.Errorf("Query has duplicate aliases %s", strings.Join(duplicateAliases, ", "))
	}

	// Check for predicate references to nonexistant events
	if q.predicate != nil {
		for _, alias := range q.predicate.usedAliases() {
			if _, ok := seenAliases[alias]; !ok {
				return fmt.Errorf("Reference to nonexistent capture %s", alias)
			}
		}
	}

	if q.window < 0 {
		return fmt.Errorf("Query has negative window duration")
	}

	return nil
}
