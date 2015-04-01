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

type PredicateResult uint8

func (r PredicateResult) And(r2 PredicateResult) PredicateResult {
	if r == r2 {
		return r
	} else if r == PredicateResultNegative || r2 == PredicateResultNegative {
		return PredicateResultNegative
	} else {
		return PredicateResultUncertain
	}
}

func (r PredicateResult) Or(r2 PredicateResult) PredicateResult {
	if r == PredicateResultPositive || r2 == PredicateResultPositive {
		return PredicateResultPositive
	} else if r == PredicateResultUncertain || r2 == PredicateResultUncertain {
		return PredicateResultUncertain
	} else {
		return PredicateResultNegative
	}
}

//go:generate stringer -type=PredicateResult

const (
	// PredicateResultPositive indicates a positive event match
	PredicateResultPositive PredicateResult = iota
	// PredicateResultNegative indicates a negative event match
	PredicateResultNegative
	// PredicateResultUncertain indicates that it cannot, with the event(s) provided, be deterimined definitively if
	// there is a match or not
	PredicateResultUncertain
)

type Predicate interface {
	Representable
	// Evaluates the predicate against the set of captured events, returning its match status.
	Evaluate(domain.CapturedEvents) PredicateResult
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

func (p *operatorPredicate) Evaluate(evs domain.CapturedEvents) PredicateResult {
	leftVal, rightVal, err := leftRightVals(evs, p.left, p.right)
	if err == ErrEventNotFound {
		return PredicateResultUncertain
	} else if err != nil {
		log.Errorf("[sase:operatorPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
		return PredicateResultNegative // Terminate this match
	}

	switch p.op {
	case opEq:
		if reflect.DeepEqual(leftVal, rightVal) {
			return PredicateResultPositive
		}
		return PredicateResultNegative

	case opNe:
		if !reflect.DeepEqual(leftVal, rightVal) {
			return PredicateResultPositive
		}
		return PredicateResultNegative

	// >, <, >=, <= only work for float64's (currently)
	case opGt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal > rightVal {
					return PredicateResultPositive
				}
				return PredicateResultNegative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare gt for non-float64s: %s", p.QueryText())
		return PredicateResultNegative // Terminate this match

	case opLt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal < rightVal {
					return PredicateResultPositive
				}
				return PredicateResultNegative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare lt for non-float64s: %s", p.QueryText())
		return PredicateResultNegative // Terminate this match

	case opGe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal >= rightVal {
					return PredicateResultPositive
				}
				return PredicateResultNegative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare ge for non-float64s: %s", p.QueryText())
		return PredicateResultNegative // Terminate this match

	case opLe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				if leftVal <= rightVal {
					return PredicateResultPositive
				}
				return PredicateResultNegative
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare le for non-float64s: %s", p.QueryText())
		return PredicateResultNegative // Terminate this match

	default:
		log.Errorf("[sase:operatorPredicate] Unhandled op %v for %s", p.op, p.QueryText())
		return PredicateResultNegative
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

func (p equivalenceTestPredicate) Evaluate(evs domain.CapturedEvents) PredicateResult {
	var (
		lastVal interface{}
		i       = 0
	)
	for alias, _ := range evs {
		if val, err := attributeLookup(fmt.Sprintf("%s.%s", alias, p)).Value(evs); err != nil {
			if err == ErrEventNotFound {
				return PredicateResultUncertain
			} else if err != nil {
				log.Errorf("[sase:equivalenceTestPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
				return PredicateResultNegative // Terminate this match
			}
		} else {
			if i > 0 && !reflect.DeepEqual(lastVal, val) {
				return PredicateResultNegative
			}
			lastVal = val
			i++
		}
	}
	return PredicateResultPositive
}

func (p equivalenceTestPredicate) QueryText() string {
	return fmt.Sprintf("[%s]", p)
}

func (p equivalenceTestPredicate) usedAliases() []string { // Not much we can do here :(
	return nil
}
