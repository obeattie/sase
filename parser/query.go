package parser

import (
	"time"
)

type Identifier string

// A selector represents an attribute selection (ie. "x.y.z")
type Selector []Identifier

type Predicate struct {
}

type Query struct {
	Events []string
	Where  []string
	Within time.Duration
}
