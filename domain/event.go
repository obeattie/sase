package domain

import (
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
