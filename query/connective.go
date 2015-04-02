package query

import (
	"strings"

	"github.com/obeattie/sase/domain"
)

// A conjunction represents an array of AND'ed predicates
type conjunction []Predicate

func (c conjunction) Evaluate(evs domain.CapturedEvents) Result {
	for _, p := range c {
		switch p.Evaluate(evs) {
		case Uncertain:
			return Uncertain
		case Negative:
			return Negative
		}
	}
	return Positive
}

func (c conjunction) QueryText() string {
	results := make([]string, len(c))
	for i, p := range c {
		results[i] = p.QueryText()
	}
	return strings.Join(results, " AND ")
}

func (c conjunction) usedAliases() []string {
	result := make([]string, 0, len(c))
	for _, p := range c {
		result = append(result, p.usedAliases()...)
	}
	return result
}

// A disunction represents an array of OR'd predicates
type disjunction []Predicate

func (d disjunction) Evaluate(evs domain.CapturedEvents) Result {
	for _, p := range d {
		switch p.Evaluate(evs) {
		case Uncertain:
			return Uncertain
		case Positive:
			return Positive
		}
	}
	return Negative
}

func (d disjunction) QueryText() string {
	results := make([]string, len(d))
	for i, p := range d {
		results[i] = p.QueryText()
	}
	return strings.Join(results, " OR ")
}

func (d disjunction) usedAliases() []string {
	result := make([]string, 0, len(d))
	for _, p := range d {
		result = append(result, p.usedAliases()...)
	}
	return result
}
