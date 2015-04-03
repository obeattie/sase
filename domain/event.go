package domain

import (
	"fmt"
	"strings"
	"time"
)

type Event interface {
	// Returns the type of the event
	Type() string
	// Returns a map of the event's attributes
	Attributes() map[string]interface{}
	// When returns the timestamp at which the event occurred
	When() time.Time
}

// CapturedEvents represents an event sequence that has been captured from an input stream (or streams) by a query,
// mapping aliases to actual events
type CapturedEvents map[string]Event

func DescribeCapturedEvents(evs CapturedEvents) string {
	results := make([]string, 0, len(evs))

	for alias, event := range evs {
		desc := fmt.Sprintf("%s:%s", alias, event.Type())
		results = append(results, desc)
	}

	return strings.Join(results, " ")
}
