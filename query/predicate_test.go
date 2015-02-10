package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obeattie/sase/domain"
)

func TestOperatorPredicate(t *testing.T) {
	e1 := &tEventImpl{
		typ: "e1",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "bar",
			"num": float64(1.5),
		},
	}
	e2 := &tEventImpl{
		typ: "e2",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "baz",
			"num": float64(1.5),
		},
	}
	e3 := &tEventImpl{
		typ: "e3",
		attrs: map[string]interface{}{
			"foo": "foobar",
			"bar": "baz",
			"num": float64(2.0),
		},
	}
	evs := domain.CapturedEvents{
		"e1": e1,
		"e2": e2,
		"e3": e3,
	}

	cases := map[op]map[[2]string]PredicateResult{
		opEq: {
			[2]string{"e1.foo", "e2.foo"}:       PredicateResultPositive,
			[2]string{"e1.foo", "e3.foo"}:       PredicateResultNegative,
			[2]string{"e2.bar", "e3.bar"}:       PredicateResultPositive,
			[2]string{"e1.bar", "e3.bar"}:       PredicateResultNegative,
			[2]string{"e1.foo", "e3.bazbazbas"}: PredicateResultNegative,  // Attribute not found
			[2]string{"e1.foo", "e5.foo"}:       PredicateResultUncertain, // Event not found
		},
		opNe: { // The inverse of opEq's cases
			[2]string{"e1.foo", "e2.foo"}:       PredicateResultNegative,
			[2]string{"e1.foo", "e3.foo"}:       PredicateResultPositive,
			[2]string{"e2.bar", "e3.bar"}:       PredicateResultNegative,
			[2]string{"e1.bar", "e3.bar"}:       PredicateResultPositive,
			[2]string{"e1.foo", "e3.bazbazbas"}: PredicateResultNegative, // Attribute not found always return false
			[2]string{"e1.foo", "e5.foo"}:       PredicateResultUncertain,
		},
		opLt: {
			[2]string{"e1.num", "e3.num"}:       PredicateResultPositive,
			[2]string{"e3.num", "e1.num"}:       PredicateResultNegative,
			[2]string{"e1.num", "e2.num"}:       PredicateResultNegative,
			[2]string{"e3.numfoo", "e1.numfoo"}: PredicateResultNegative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       PredicateResultNegative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       PredicateResultNegative,
		},
		opGt: {
			[2]string{"e1.num", "e3.num"}:       PredicateResultNegative,
			[2]string{"e3.num", "e1.num"}:       PredicateResultPositive,
			[2]string{"e1.num", "e2.num"}:       PredicateResultNegative,
			[2]string{"e3.numfoo", "e1.numfoo"}: PredicateResultNegative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       PredicateResultNegative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       PredicateResultNegative,
		},
		opLe: {
			[2]string{"e1.num", "e3.num"}:       PredicateResultPositive,
			[2]string{"e3.num", "e1.num"}:       PredicateResultNegative,
			[2]string{"e1.num", "e2.num"}:       PredicateResultPositive,
			[2]string{"e3.numfoo", "e1.numfoo"}: PredicateResultNegative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       PredicateResultNegative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       PredicateResultNegative,
		},
		opGe: {
			[2]string{"e1.num", "e3.num"}:       PredicateResultNegative,
			[2]string{"e3.num", "e1.num"}:       PredicateResultPositive,
			[2]string{"e1.num", "e2.num"}:       PredicateResultPositive,
			[2]string{"e3.numfoo", "e1.numfoo"}: PredicateResultNegative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       PredicateResultNegative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       PredicateResultNegative,
		},
	}

	for op, opCases := range cases {
		for operands, expectedResult := range opCases {
			impl := &operatorPredicate{
				left:  attributeLookup(operands[0]),
				right: attributeLookup(operands[1]),
				op:    op,
			}

			assert.Equal(t, expectedResult, (impl.Evaluate(evs)), fmt.Sprintf("Incorrect result for \"%s\"",
				impl.QueryText()))
		}
	}
}
