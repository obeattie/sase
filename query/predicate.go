package query

import (
	"bytes"
	"fmt"
	"reflect"

	log "github.com/cihub/seelog"

	"github.com/obeattie/sase/domain"
)

func leftRightVals(evs domain.CapturedEvents, left, right value) (interface{}, interface{}, error) {
	if left == nil || right == nil {
		return nil, nil, fmt.Errorf("Left and right must not be nil")
	} else if leftVal, err := left.Value(evs); err != nil {
		return nil, nil, err
	} else if rightVal, err := right.Value(evs); err != nil {
		return nil, nil, err
	} else {
		return leftVal, rightVal, err
	}
}

type Result uint8

func (r Result) And(r2 Result) Result {
	if r == r2 {
		return r
	} else if r == Negative || r2 == Negative {
		return Negative
	} else {
		return Uncertain
	}
}

func (r Result) Or(r2 Result) Result {
	if r == Positive || r2 == Positive {
		return Positive
	} else if r == Uncertain || r2 == Uncertain {
		return Uncertain
	} else {
		return Negative
	}
}

//go:generate stringer -type=Result

const (
	// Positive indicates a positive event match
	Positive Result = iota
	// Negative indicates a negative event match
	Negative
	// Uncertain indicates that it cannot, with the event(s) provided, be deterimined definitively if
	// there is a match or not
	Uncertain
)

type Predicate interface {
	Representable
	// Evaluates the predicate against the set of captured events, returning its match status.
	Evaluate(domain.CapturedEvents) Result
	// usedAliases returns the events aliases which are consulted during evaluation
	usedAliases() []string
}

type op uint8

const (
	opEq op = iota // equal to (==)
	opNe           // not equal to (!=)
	opGt           // greater than (>)
	opLt           // less than (<)
	opGe           // greater than or equal to (>=)
	opLe           // less than or equal to (<=)
)

// An operatorPredicate evaluates an operator between two values
type operatorPredicate struct {
	left  value
	right value
	op    op
}

func (p *operatorPredicate) Evaluate(evs domain.CapturedEvents) Result {
	leftVal, rightVal, err := leftRightVals(evs, p.left, p.right)
	if err == ErrEventNotFound {
		return Uncertain
	} else if err != nil {
		log.Errorf("[sase:operatorPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
		return Negative // Terminate this match
	}

	switch p.op {
	case opEq:
		if reflect.DeepEqual(leftVal, rightVal) {
			return Positive
		}
		return Negative

	case opNe:
		if !reflect.DeepEqual(leftVal, rightVal) {
			return Positive
		}
		return Negative

	// >, <, >=, <= only work for float64's (currently)
	case opGt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal > rightVal {
					return Positive
				}
				return Negative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare gt for non-float64s: %s", p.QueryText())
		return Negative // Terminate this match

	case opLt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal < rightVal {
					return Positive
				}
				return Negative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare lt for non-float64s: %s", p.QueryText())
		return Negative // Terminate this match

	case opGe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal >= rightVal {
					return Positive
				}
				return Negative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare ge for non-float64s: %s", p.QueryText())
		return Negative // Terminate this match

	case opLe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal <= rightVal {
					return Positive
				}
				return Negative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare le for non-float64s: %s", p.QueryText())
		return Negative // Terminate this match

	default:
		log.Errorf("[sase:operatorPredicate] Unhandled op %v for %s", p.op, p.QueryText())
		return Negative
	}
}

func (p *operatorPredicate) QueryText() string {
	buf := new(bytes.Buffer)
	if p.left != nil {
		buf.WriteString(p.left.QueryText())
	}
	buf.WriteRune(' ')
	switch p.op {
	case opEq:
		buf.WriteString("==")
	case opNe:
		buf.WriteString("!=")
	case opGt:
		buf.WriteRune('>')
	case opLt:
		buf.WriteRune('<')
	case opGe:
		buf.WriteString(">=")
	case opLe:
		buf.WriteString("<=")
	}
	if p.right != nil {
		buf.WriteRune(' ')
		buf.WriteString(p.right.QueryText())
	}
	return buf.String()
}

func (p *operatorPredicate) usedAliases() []string {
	result := make([]string, 0)
	if p.left != nil {
		result = append(result, p.left.usedAliases()...)
	}
	if p.right != nil {
		result = append(result, p.right.usedAliases()...)
	}
	return result
}

type equivalenceTestPredicate string // Holds the equivalence key path

func (p equivalenceTestPredicate) Evaluate(evs domain.CapturedEvents) Result {
	var (
		lastVal interface{}
		i       = 0
	)
	for alias, _ := range evs {
		if val, err := attributeLookup(fmt.Sprintf("%s.%s", alias, p)).Value(evs); err != nil {
			if err == ErrEventNotFound {
				return Uncertain
			} else if err != nil {
				log.Errorf("[sase:equivalenceTestPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
				return Negative // Terminate this match
			}
		} else {
			if i > 0 && !reflect.DeepEqual(lastVal, val) {
				return Negative
			}
			lastVal = val
			i++
		}
	}
	return Positive
}

func (p equivalenceTestPredicate) QueryText() string {
	return fmt.Sprintf("[%s]", p)
}

func (p equivalenceTestPredicate) usedAliases() []string { // Not much we can do here :(
	return nil
}
