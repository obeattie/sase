package domain

type Event interface {
	// Returns the type of the event
	Type() string
	// Returns a map of the event's attributes
	Attributes() map[string]interface{}
}

// CapturedEvents represents an event sequence that has been captured from an input stream (or streams) by a query.
type CapturedEvents map[string]Event
