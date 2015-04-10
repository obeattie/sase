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
   * have a type "t{{n}}", where n is as given above
   * have all the attributes of the previous event, with an additional attribute (randomly chosen from a map of
     candidates).
   * have a unique attribute "ue{{n}}" (=1), where n is as given above
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
		attrs := map[string]interface{}{}
		for k, v := range lastAttrs {
			attrs[k] = v
		}
		attrs["u"+key] = 1
		result[i] = &tEventImpl{
			typ:   fmt.Sprintf("t%d", i),
			attrs: attrs,
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
			var newResult []domain.CapturedEvents
		acceptingLoop:
			for _, stack := range result {
				newResult = append(newResult, stack)
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
				if q.Evaluate(newStack) == Invalid {
					newResult = newResult[:len(newResult)-1]
				} else {
					newResult = append(newResult, newStack)
				}
			}

			// Create a new (virgin) stack with this capture and no others
			result = append(newResult, domain.CapturedEvents{alias: e})
		}
	}
	return result
}

func TestE2E(t *testing.T) {
	events := genEvents(100)
	queries := map[string]bool{ // true = one (and only one) positive match
		`EVENT t0 e0`:                                                   true,  // No predicate should match
		`EVENT t0 e0 WHERE e0.foobar == e0.foobaz`:                      false, // Nonexistent events
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 == e1.e0`:                  true,
		`EVENT SEQ(t0 e0, t1 e1) WHERE e0.e0 != e1.e0`:                  false,
		`EVENT SEQ(t0 e0, t1000 e1) WHERE e0.e0 == e1.e0`:               false, // Incomplete events
		`EVENT SEQ(t0 e0, t1 e1, t1000 e2) WHERE e0.e0 != e1.e0`:        false, // Incomplete events but known to not match
		`EVENT SEQ(t1 e1, t0 e0)`:                                       false, // SEQ out of order
		`EVENT SEQ(t0 e0, !(t1 e1), t2 e2)`:                             false, // Negated event in stream
		`EVENT SEQ(t0 e0, !(foo bar), t1 e1)`:                           true,  // Negated event not in stream
		`EVENT SEQ(t0 e0, !(foo bar), !(baz boo), t2 e2)`:               true,  // Negated events not in stream
		`EVENT SEQ(t0 e0, !(foo bar), !(t1 e1), !(baz boo), t2 e2)`:     false, // One negated event in stream
		`EVENT SEQ(t0 e0, !(t1 e1_n), t1 e1)`:                           true,  // Only 1x e1 should be encountered
		`EVENT SEQ(t0 e0, !(t1 e1_n), !(t1 e1_n1), t1 e1)`:              true,
		`EVENT SEQ(t0 e0, !(t1 e1_n1), !(t0 e0_n), !(t1 e1_n2), t1 e1)`: true,
		`EVENT ANY(t0 e0, foo foo)`:                                     true,
		`EVENT ANY(foo foo, bar bar, baz baz)`:                          false,
		`EVENT SEQ(t0 e0, ANY(foo foo, t1 e1))`:                         true,

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
		// …on optional events
		`EVENT SEQ(t0 e0, ANY(t1 e1, foo bar)) WHERE [string]`:              true,  // foo not present
		`EVENT SEQ(t0 e0, ANY(t1 e1)) WHERE [ue0]`:                          false, // ue0 not present on e1
		`EVENT SEQ(t0 e0, !(foo bar), t1 e1) WHERE [string]`:                true,  // foo not present
		`EVENT SEQ(t0 e0, !(foo bar), t1 e1) WHERE e0.string == bar.string`: true,

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
		found := false
		for _, stack := range stacks {
			r := q.Evaluate(stack)
			t.Logf("    [%s]: %s", domain.DescribeCapturedEvents(stack), r.String())
			switch r {
			case Positive:
				require.False(t, found, "More than 1 match found")
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

	testEventSequence(t, queryText, events, true)
}

func TestE2ENegatedSequence(t *testing.T) {
	events := []domain.Event{
		&tEventImpl{
			typ: "t0",
			ts:  time.Date(2014, 1, 1, 0, 0, 20, 0, time.UTC),
		},
		&tEventImpl{
			typ: "t0",
			ts:  time.Date(2014, 1, 1, 0, 0, 20, 1, time.UTC),
			attrs: map[string]interface{}{
				"attr": "val",
			},
		},
	}

	queryText := `EVENT SEQ(t0 e0, !(t0 e1), t0 e2) WHERE e1.attr == "val"`

	testEventSequence(t, queryText, events, false)
}

func testEventSequence(t *testing.T, queryText string, events []domain.Event, expectMatch bool) {
	q, err := Parse(queryText)
	require.NoError(t, err, fmt.Sprintf("Unexpected error parsing query \"%s\"", queryText))

	found := false
	stacks := genStacks(events, q)
	for _, stack := range stacks {
		r := q.Evaluate(stack)
		t.Logf("[%s]: %s", domain.DescribeCapturedEvents(stack), r.String())
		if r == Positive {
			require.False(t, found, "More than 1 match found")
			found = true
		}
	}

	require.Equal(t, expectMatch, found, fmt.Sprintf("Unexpected match for \"%s\"", queryText))
}
