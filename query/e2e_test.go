package query

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/obeattie/sase/domain"
)

/* GenEvents returns a slice of events of the given length. These events:

   * Will have aliases "e{{n}}", where n is its index in the slice (0-indexed)
   * Will have a type "t{{n}}", where n is the same value as above
   * Will have all the attributes of the previous event, with an additional attribute (randomly chosen from a map of
     candidates).
   * Will have static attributes "map", "array", "decimal", "string", and "stringdecimal"
*/
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

	var (
		lastAttrs = map[string]interface{}{
			"array":         []string{"1", "2", "3"},
			"decimal":       100.0,
			"string":        "astring",
			"stringdecimal": "100.0",
			"map": map[string]interface{}{
				"key": "value",
			},
		}
		lastTs = time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)
	)
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
		lastTs = lastTs.Add(1 * time.Minute)
		lastAttrs = newAttrs(key)
		result[i] = &tEventImpl{
			typ:   fmt.Sprintf("t%d", i),
			attrs: lastAttrs,
			ts:    lastTs,
		}
	}
	return result
}

func TestE2E(t *testing.T) {
	events := genEvents(100)
	queries := map[string]PredicateResult{
		`EVENT t0 e0`:                                            PredicateResultPositive, // No predicate should match
		`EVENT t0 e0 WHERE e0.foobar == e0.foobaz`:               PredicateResultNegative, // Nonexistent events
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 == e1.e0`:           PredicateResultPositive,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 != e1.e0`:           PredicateResultNegative,
		`EVENT SEQ(t0 e0, t1000 e1) WHERE e0.e0 == e1.e0`:        PredicateResultUncertain, // Incomplete events
		`EVENT SEQ(t0 e0, t1 e1, t1000 e2) WHERE e0.e0 != e1.e0`: PredicateResultNegative,  // Incomplete events but known to not match
		`EVENT SEQ(t1 e1, t0 e0)`:                                PredicateResultNegative,  // SEQ out of order
		`EVENT SEQ(t0 e0, !(t1 e1), t2 e2)`:                      PredicateResultNegative,  // Negated event is in stream

		// Attribute tests
		`EVENT t0 e0 WHERE e0.string == "astring"`:                             PredicateResultPositive,
		`EVENT t0 e0 WHERE e0.nonexistent == 1`:                                PredicateResultNegative,
		`EVENT t0 e0 WHERE e0.stringdecimal == 1.0`:                            PredicateResultNegative, // Wrong type
		`EVENT t0 e0 WHERE e0.decimal == 100`:                                  PredicateResultPositive,
		`EVENT t0 e0 WHERE e0.decimal == 100.0`:                                PredicateResultPositive,
		`EVENT t0 e0 WHERE e0.decimal == 100.00000000`:                         PredicateResultPositive,
		`EVENT t0 e0 WHERE e0.decimal == 0000100.00`:                           PredicateResultPositive, // Different padding
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.decimal == e1.decimal`:               PredicateResultPositive,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.array == e1.array`:                   PredicateResultPositive,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.string == e1.string`:                 PredicateResultPositive,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.stringdecimal == e1.stringdecimal`:   PredicateResultPositive,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.map == e1.map`:                       PredicateResultPositive, // Static attibutes
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.decimal == e1.decimal`:             PredicateResultPositive,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.array == e1.array`:                 PredicateResultPositive,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.string == e1.string`:               PredicateResultPositive,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.stringdecimal == e1.stringdecimal`: PredicateResultPositive,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.map == e1.map`:                     PredicateResultPositive,
		`EVENT t0 e0 WHERE [string]`:                                           PredicateResultPositive,
		`EVENT SEQ(t0 e0) WHERE [string]`:                                      PredicateResultPositive,  // Equivalence test
		`EVENT SEQ(t0 e0, t1 e1, t2 e2) WHERE [string]`:                        PredicateResultPositive,  // Equivalence test
		`EVENT SEQ(t0 e0, t1 e1) WHERE [e1]`:                                   PredicateResultNegative,  // e1 attr will be missing from e0
		`EVENT SEQ(t0 e0, e1 e1, t1000 e1000) WHERE [string]`:                  PredicateResultUncertain, // No t1000
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
		assert.Equal(t, expectedResult.String(), result.String(), fmt.Sprintf("Incorrect result for \"%s\"", queryText))
	}
}
