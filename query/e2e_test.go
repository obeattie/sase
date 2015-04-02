package query

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/obeattie/sase/domain"
)

/* GenEvents returns a slice of events of the given length. These events:

   * have aliases "e{{n}}", where n is its index in the slice (0-indexed)
   * have a type "t{{n}}", where n is the same value as above
   * have all the attributes of the previous event, with an additional attribute (randomly chosen from a map of
     candidates).
   * have static attributes "map", "array", "decimal", "string", and "stringdecimal"
   * are sepearated by 1 minute intervals
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

func genStacks(evs []domain.Event, q *Query) []domain.CapturedEvents {
	result := make([]domain.CapturedEvents, 0)
	for _, e := range evs {
		for _, alias := range q.CaptureAliases(e) {
			// Add to any stacks that accept this event
			newStacks := make([]domain.CapturedEvents, 0)
		acceptingLoop:
			for _, stack := range result {
				if _, ok := stack[alias]; ok {
					continue acceptingLoop
				}
				for _, ev := range stack {
					if ev == e { // Can't insert a duplicate
						continue acceptingLoop
					}
				}
				newStack := domain.CapturedEvents{}
				for k, v := range stack {
					newStack[k] = v
				}
				newStack[alias] = e
				newStacks = append(newStacks, newStack)
			}
			result = append(result, newStacks...)

			// Create a new (virgin) stack with this capture and no others
			result = append(result, domain.CapturedEvents{alias: e})
		}
	}
	return result
}

func TestE2E(t *testing.T) {
	events := genEvents(100)
	queries := map[string]bool{ // true = one positive match
		`EVENT t0 e0`:                                            true,  // No predicate should match
		`EVENT t0 e0 WHERE e0.foobar == e0.foobaz`:               false, // Nonexistent events
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 == e1.e0`:           true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 != e1.e0`:           false,
		`EVENT SEQ(t0 e0, t1000 e1) WHERE e0.e0 == e1.e0`:        false, // Incomplete events
		`EVENT SEQ(t0 e0, t1 e1, t1000 e2) WHERE e0.e0 != e1.e0`: false, // Incomplete events but known to not match
		`EVENT SEQ(t1 e1, t0 e0)`:                                false, // SEQ out of order
		`EVENT SEQ(t0 e0, !(t1 e1), t2 e2)`:                      false, // Negated event is in stream

		// Attribute tests
		`EVENT t0 e0 WHERE e0.string == "astring"`:                             true,
		`EVENT t0 e0 WHERE e0.nonexistent == 1`:                                false,
		`EVENT t0 e0 WHERE e0.stringdecimal == 1.0`:                            false, // Wrong type
		`EVENT t0 e0 WHERE e0.decimal == 100`:                                  true,
		`EVENT t0 e0 WHERE e0.decimal == 100.0`:                                true,
		`EVENT t0 e0 WHERE e0.decimal == 100.00000000`:                         true,
		`EVENT t0 e0 WHERE e0.decimal == 0000100.00`:                           true, // Different padding
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.decimal == e1.decimal`:               true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.array == e1.array`:                   true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.string == e1.string`:                 true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.stringdecimal == e1.stringdecimal`:   true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.map == e1.map`:                       true, // Static attibutes
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.decimal == e1.decimal`:             true,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.array == e1.array`:                 true,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.string == e1.string`:               true,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.stringdecimal == e1.stringdecimal`: true,
		`EVENT SEQ(t10 e0, t99 e1) WHERE e0.map == e1.map`:                     true,
		`EVENT t0 e0 WHERE [string]`:                                           true,
		`EVENT SEQ(t0 e0) WHERE [string]`:                                      true,  // Equivalence test
		`EVENT SEQ(t0 e0, t1 e1, t2 e2) WHERE [string]`:                        true,  // Equivalence test
		`EVENT SEQ(t0 e0, t1 e1) WHERE [e1]`:                                   false, // e1 attr will be missing from e0
		`EVENT SEQ(t0 e0, t1 e1, t1000 e1000) WHERE [string]`:                  false, // No t1000
		`EVENT SEQ(t0 e0, t1 e1) WHERE [string] AND e0.decimal == e1.decimal`:  true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE [e1] AND e0.decimal == e1.decimal`:      false,
		`EVENT SEQ(t0 e0, t1 e1) WHERE [e1] OR e0.decimal == e1.decimal`:       true,

		// Window tests
		`EVENT SEQ(t0 e0, t10 e10) WITHIN 1m`:    false, // e0 and e10 are 10 minutes apart
		`EVENT SEQ(t0 e0, t10 e10) WITHIN 10m`:   true,
		`EVENT SEQ(t0 e0, t10 e10) WITHIN 1h`:    true,
		`EVENT SEQ(t0 e0, t10 e10) WITHIN 9m59s`: false,
	}

	for queryText, expectedResult := range queries {
		t.Log(queryText)
		q, err := Parse(queryText)
		require.NoError(t, err, fmt.Sprintf("Unexpected error parsing query \"%s\"", queryText))

		stacks := genStacks(events, q)
		t.Logf("Generated %d stacks for \"%s\"", len(stacks), queryText)
		found := false
		for _, stack := range stacks {
			if q.Evaluate(stack) == PredicateResultPositive {
				require.False(t, found, "More than 1 match found")
				t.Logf("Found positive match with stack %+v", stack)
				found = true
			}
		}

		require.Equal(t, expectedResult, found, "Incorrect match for \"%s\"", queryText)
	}
}

func TestE2EDuplicateTypes(t *testing.T) {
	events := []domain.Event{
		&tEventImpl{
			typ: "t0",
			ts:  time.Date(2014, 1, 1, 0, 0, 20, 0, time.UTC),
			attrs: map[string]interface{}{
				"attr": "val",
			}},
		&tEventImpl{
			typ: "t0",
			ts:  time.Date(2014, 1, 1, 0, 0, 10, 0, time.UTC),
			attrs: map[string]interface{}{
				"attr": "val",
			}}}

	queryText := `EVENT SEQ(t0 e0, t0 e1) WHERE [attr]`
	q, err := Parse(queryText)
	require.NoError(t, err, fmt.Sprintf("Unexpected error parsing query \"%s\"", queryText))

	found := false
	stacks := genStacks(events, q)
	for _, stack := range stacks {
		t.Logf("Stack: %v", stack)
		if q.Evaluate(stack) == PredicateResultPositive {
			require.False(t, found, "More than 1 match found")
			t.Logf("Found positive match with stack %+v", stack)
			found = true
		}
	}

	require.True(t, found, fmt.Sprintf("No positive match found for \"%s\"", queryText))
}
