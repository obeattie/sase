package domain

type Event interface {
	// Returns the type of the event
	Type() string
	// Returns a map of the event's attributes
	Attributes() map[string]interface{}
}
