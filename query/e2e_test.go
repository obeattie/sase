package query

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/obeattie/sase/domain"
)

func genEvents(length int) []domain.Event {
	result := make([]domain.Event, length)
	rand.Seed(time.Now().UnixNano())

	candidates := [...]interface{}{
		map[string]interface{}{
			"1": 2,
		},
		[]string{"1", "2", "3"},
		20.0,
		"string",
	}

	var lastAttrs map[string]interface{}
	newAttrs := func(key string) map[string]interface{} {
		ret := make(map[string]interface{}, len(lastAttrs))
		for k, v := range lastAttrs {
			ret[k] = v
		}
		ret[key] = candidates[rand.Intn(len(candidates)-1)]
		return ret
	}

	for i := 0; i < length; i++ {
		key := fmt.Sprintf("e%d", i)
		lastAttrs = newAttrs(key)
		result[i] = &tEventImpl{
			typ:   fmt.Sprintf("t%d", i),
			attrs: lastAttrs,
		}
	}
	return result
}

func TestE2E(t *testing.T) {
	events := genEvents(100)
	tr, fa := true, false
	queries := map[string]*bool{
		"EVENT t0 e0":                                  &tr, // No predicate should match
		"EVENT t0 e0 WHERE e0.foobar == e0.foobaz":     &fa, // Nonexistent events
		"EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 == e1.e0": &tr,
		"EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 != e1.e0": &fa,
	}

	for queryText, expectedResult := range queries {
		q, err := Parse(queryText)
		assert.NoError(t, err, fmt.Sprintf("Unexpected error parsing query \"%s\"", queryText))

		capturedEvents := make(map[string]domain.Event)
		for _, e := range events {
			if alias := q.ShouldCapture(e); alias != "" {
				capturedEvents[alias] = e
			}
		}
		assert.NotEmpty(t, capturedEvents, "No events captured for \"%s\"", queryText)

		result := q.Evaluate(capturedEvents)
		if expectedResult == nil {
			assert.Nil(t, result, fmt.Sprintf("Incorrect capture value for \"%s\"", queryText))
		} else {
			assert.Equal(t, *expectedResult, *result, fmt.Sprintf("Incorrect capture value for \"%s\"", queryText))
		}
	}
}
