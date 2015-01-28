package query

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obeattie/sase/domain"
)

func TestEqPredicate(t *testing.T) {
	e1 := &tEventImpl{
		typ: "e1",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "bar",
		},
	}
	e2 := &tEventImpl{
		typ: "e2",
		attrs: map[string]interface{}{
			"foo": "bar",
			"bar": "baz",
		},
	}
	e3 := &tEventImpl{
		typ: "e3",
		attrs: map[string]interface{}{
			"foo": "foobar",
			"bar": "baz",
		},
	}
	evs := domain.CapturedEvents{
		"e1": e1,
		"e2": e2,
		"e3": e3,
	}

	tr, fa := true, false
	cases := map[[2]string]*bool{
		[2]string{"e1.foo", "e2.foo"}:       &tr,
		[2]string{"e1.foo", "e3.foo"}:       &fa,
		[2]string{"e2.bar", "e3.bar"}:       &tr,
		[2]string{"e1.bar", "e3.bar"}:       &fa,
		[2]string{"e1.foo", "e3.bazbazbas"}: &fa, // Attribute not found
		[2]string{"e1.foo", "e5.foo"}:       nil, // Event not found
	}

	for operands, result := range cases {
		impl := &eqPredicate{
			left:  attributeLookup(operands[0]),
			right: attributeLookup(operands[1]),
		}
		assert.Equal(t, result, impl.Evaluate(evs))
	}
}
