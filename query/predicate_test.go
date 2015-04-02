package query

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/obeattie/sase/domain"
)

func TestPredicateResultAnd(t *testing.T) {
	expectations := map[[2]PredicateResult]PredicateResult{
		[2]PredicateResult{PredicateResultNegative, PredicateResultNegative}:   PredicateResultNegative,
		[2]PredicateResult{PredicateResultNegative, PredicateResultPositive}:   PredicateResultNegative,
		[2]PredicateResult{PredicateResultNegative, PredicateResultUncertain}:  PredicateResultNegative,
		[2]PredicateResult{PredicateResultPositive, PredicateResultNegative}:   PredicateResultNegative,
		[2]PredicateResult{PredicateResultPositive, PredicateResultPositive}:   PredicateResultPositive,
		[2]PredicateResult{PredicateResultPositive, PredicateResultUncertain}:  PredicateResultUncertain,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultNegative}:  PredicateResultNegative,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultPositive}:  PredicateResultUncertain,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultUncertain}: PredicateResultUncertain,
	}

	for inputs, expected := range expectations {
		actual := inputs[0].And(inputs[1])
		require.Equal(t, expected.String(), actual.String(), fmt.Sprintf("%s AND %s", inputs[0].String(),
			inputs[1].String()))
	}
}

func TestPredicateResultOr(t *testing.T) {
	expectations := map[[2]PredicateResult]PredicateResult{
		[2]PredicateResult{PredicateResultNegative, PredicateResultNegative}:   PredicateResultNegative,
		[2]PredicateResult{PredicateResultNegative, PredicateResultPositive}:   PredicateResultPositive,
		[2]PredicateResult{PredicateResultNegative, PredicateResultUncertain}:  PredicateResultUncertain,
		[2]PredicateResult{PredicateResultPositive, PredicateResultNegative}:   PredicateResultPositive,
		[2]PredicateResult{PredicateResultPositive, PredicateResultPositive}:   PredicateResultPositive,
		[2]PredicateResult{PredicateResultPositive, PredicateResultUncertain}:  PredicateResultPositive,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultNegative}:  PredicateResultUncertain,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultPositive}:  PredicateResultPositive,
		[2]PredicateResult{PredicateResultUncertain, PredicateResultUncertain}: PredicateResultUncertain,
	}

	for inputs, expected := range expectations {
		actual := inputs[0].Or(inputs[1])
		require.Equal(t, expected.String(), actual.String(), fmt.Sprintf("%s OR %s", inputs[0].String(),
			inputs[1].String()))
	}
}

func TestOperatorPredicate(t *testing.T) {
	e1 := &tEventImpl{
		typ: "e1",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "bar",
			"num": float64(1.5),
		},
		ts: time.Now(),
	}
	e2 := &tEventImpl{
		typ: "e2",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "baz",
			"num": float64(1.5),
		},
		ts: time.Now(),
	}
	e3 := &tEventImpl{
		typ: "e3",
		attrs: map[string]interface{}{
			"foo": "foobar",
			"bar": "baz",
			"num": float64(2.0),
		},
		ts: time.Now(),
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

			require.Equal(t, expectedResult, (impl.Evaluate(evs)), fmt.Sprintf("Incorrect result for \"%s\"",
				impl.QueryText()))
		}
	}
}
