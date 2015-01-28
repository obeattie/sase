package query

import (
	"fmt"
	"reflect"

	log "github.com/cihub/seelog"
)

func leftRightVals(evs CapturedEvents, left value, right value) (interface{}, interface{}, error) {
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
	Evaluate(CapturedEvents) *bool
}

// An eqPredicate evaluates equality between the two passed events
type eqPredicate struct {
	left  value
	right value
}

func (p *eqPredicate) Evaluate(evs CapturedEvents) *bool {
	leftVal, rightVal, err := leftRightVals(evs, p.left, p.right)
	if err == ErrEventNotFound {
		return nil
	} else if err != nil {
		log.Errorf("[sase:eqPredicate] Could not evaluate %s left/right: %s", p.QueryText(), err.Error())
		ret := false
		return &ret // Terminate this match
	}

	equal := reflect.DeepEqual(leftVal, rightVal)
	return &equal
}

func (p *eqPredicate) QueryText() string {
	if p.left != nil && p.right != nil {
		return fmt.Sprintf("%s == %s", p.left.QueryText(), p.right.QueryText())
	} else {
		return ""
	}
}
