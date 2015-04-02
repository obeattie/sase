package query

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/obeattie/sase/domain"
)

func TestAnd(t *testing.T) {
	expectations := map[[2]Result]Result{
		[2]Result{Negative, Negative}:   Negative,
		[2]Result{Negative, Positive}:   Negative,
		[2]Result{Negative, Uncertain}:  Negative,
		[2]Result{Positive, Negative}:   Negative,
		[2]Result{Positive, Positive}:   Positive,
		[2]Result{Positive, Uncertain}:  Uncertain,
		[2]Result{Uncertain, Negative}:  Negative,
		[2]Result{Uncertain, Positive}:  Uncertain,
		[2]Result{Uncertain, Uncertain}: Uncertain,
	}

	for inputs, expected := range expectations {
		actual := inputs[0].And(inputs[1])
		require.Equal(t, expected.String(), actual.String(), fmt.Sprintf("%s AND %s", inputs[0].String(),
			inputs[1].String()))
	}
}

func TestOr(t *testing.T) {
	expectations := map[[2]Result]Result{
		[2]Result{Negative, Negative}:   Negative,
		[2]Result{Negative, Positive}:   Positive,
		[2]Result{Negative, Uncertain}:  Uncertain,
		[2]Result{Positive, Negative}:   Positive,
		[2]Result{Positive, Positive}:   Positive,
		[2]Result{Positive, Uncertain}:  Positive,
		[2]Result{Uncertain, Negative}:  Uncertain,
		[2]Result{Uncertain, Positive}:  Positive,
		[2]Result{Uncertain, Uncertain}: Uncertain,
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

	cases := map[op]map[[2]string]Result{
		opEq: {
			[2]string{"e1.foo", "e2.foo"}:       Positive,
			[2]string{"e1.foo", "e3.foo"}:       Negative,
			[2]string{"e2.bar", "e3.bar"}:       Positive,
			[2]string{"e1.bar", "e3.bar"}:       Negative,
			[2]string{"e1.foo", "e3.bazbazbas"}: Negative,  // Attribute not found
			[2]string{"e1.foo", "e5.foo"}:       Uncertain, // Event not found
		},
		opNe: { // The inverse of opEq's cases
			[2]string{"e1.foo", "e2.foo"}:       Negative,
			[2]string{"e1.foo", "e3.foo"}:       Positive,
			[2]string{"e2.bar", "e3.bar"}:       Negative,
			[2]string{"e1.bar", "e3.bar"}:       Positive,
			[2]string{"e1.foo", "e3.bazbazbas"}: Negative, // Attribute not found always return false
			[2]string{"e1.foo", "e5.foo"}:       Uncertain,
		},
		opLt: {
			[2]string{"e1.num", "e3.num"}:       Positive,
			[2]string{"e3.num", "e1.num"}:       Negative,
			[2]string{"e1.num", "e2.num"}:       Negative,
			[2]string{"e3.numfoo", "e1.numfoo"}: Negative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       Negative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       Negative,
		},
		opGt: {
			[2]string{"e1.num", "e3.num"}:       Negative,
			[2]string{"e3.num", "e1.num"}:       Positive,
			[2]string{"e1.num", "e2.num"}:       Negative,
			[2]string{"e3.numfoo", "e1.numfoo"}: Negative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       Negative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       Negative,
		},
		opLe: {
			[2]string{"e1.num", "e3.num"}:       Positive,
			[2]string{"e3.num", "e1.num"}:       Negative,
			[2]string{"e1.num", "e2.num"}:       Positive,
			[2]string{"e3.numfoo", "e1.numfoo"}: Negative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       Negative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       Negative,
		},
		opGe: {
			[2]string{"e1.num", "e3.num"}:       Negative,
			[2]string{"e3.num", "e1.num"}:       Positive,
			[2]string{"e1.num", "e2.num"}:       Positive,
			[2]string{"e3.numfoo", "e1.numfoo"}: Negative, // Attribute not found always returns false
			[2]string{"e1.foo", "e3.num"}:       Negative, // Wrong attribute type always returns false
			[2]string{"e3.foo", "e1.num"}:       Negative,
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
