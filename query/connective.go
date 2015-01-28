package query

import (
	"strings"

	"github.com/obeattie/sase/domain"
)

// A conjunction represents an array of AND'ed predicates
type conjunction []Predicate

func (c conjunction) Evaluate(evs domain.CapturedEvents) *bool {
	result := true
	for _, p := range c {
		r := p.Evaluate(evs)
		if r == nil {
			return nil
		}
		result = result || *r
		if !result {
			break
		}
	}
	return &result
}

func (c conjunction) QueryText() string {
	results := make([]string, len(c))
	for i, p := range c {
		results[i] = p.QueryText()
	}
	return strings.Join(results, " AND ")
}

// A disunction represents an array of OR'd predicates
type disjunction []Predicate

func (d disjunction) Evaluate(evs domain.CapturedEvents) *bool {
	result := false
	hasNil := false
	for _, p := range d {
		r := p.Evaluate(evs)
		if r == nil {
			hasNil = true
			continue
		}
		result = result || *r
	}
	if !result && hasNil {
		return nil
	}
	return &result
}

func (d disjunction) QueryText() string {
	results := make([]string, len(d))
	for i, p := range d {
		results[i] = p.QueryText()
	}
	return strings.Join(results, " OR ")
}
