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

type Predicate interface {
	Representable
	// Evaluates the predicate against the set of captured events, returning its match status. It returns a bool pointer
	// so that it may return nil, which means that with the current event set, it's not possible to evaluate the
	// predicate (ie. it refers to events which have not yet been captured).
	// TODO: Handle nil values higher up to always terminate a candidate when its event sequence is known to be complete.
	Evaluate(domain.CapturedEvents) *bool
	// usedAliases returns the events aliases which are consulted during evaluation
	usedAliases() []string
}

type op uint8

const (
	opEq op = iota
	opNe
	opGt
	opLt
	opGe
	opLe
)

// An operatorPredicate evaluates an operator between two values
type operatorPredicate struct {
	left  value
	right value
	op    op
}

func (p *operatorPredicate) Evaluate(evs domain.CapturedEvents) *bool {
	leftVal, rightVal, err := leftRightVals(evs, p.left, p.right)
	f := false

	if err == ErrEventNotFound {
		return nil
	} else if err != nil {
		log.Errorf("[sase:operatorPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
		return &f // Terminate this match
	}

	switch p.op {
	case opEq:
		result := reflect.DeepEqual(leftVal, rightVal)
		return &result

	case opNe:
		result := !reflect.DeepEqual(leftVal, rightVal)
		return &result

	// >, <, >=, <= only work for float64's (currently)
	case opGt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				result := leftVal > rightVal
				return &result
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare gt for non-float64s: %s", p.QueryText())
		return &f // Terminate this match

	case opLt:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				result := leftVal < rightVal
				return &result
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare lt for non-float64s: %s", p.QueryText())
		return &f // Terminate this match

	case opGe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				result := leftVal >= rightVal
				return &result
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare ge for non-float64s: %s", p.QueryText())
		return &f // Terminate this match

	case opLe:
		if leftVal, ok := leftVal.(float64); ok {
			if rightVal, ok := rightVal.(float64); ok {
				result := leftVal <= rightVal
				return &result
			}
		}
		log.Errorf("[sase:operatorPredicate] Could not compare le for non-float64s: %s", p.QueryText())
		return &f // Terminate this match

	default:
		log.Errorf("[sase:operatorPredicate] Unhandled op %v for %s", p.op, p.QueryText())
		return &f
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
