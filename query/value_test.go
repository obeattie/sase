package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeLookup(t *testing.T) {
	e := &tEventImpl{
		typ: "a",
		attrs: map[string]interface{}{
			"str": "str",
			"map": map[string]interface{}{
				"str": "str1",
				"int": 2,
				"map": map[string]string{
					"str": "str2",
				},
				"slice": []string{"1", "2", "3"},
			},
		}}

	lookups := map[string]interface{}{
		"ev1.str":         "str",
		"ev1.map":         e.attrs["map"],
		"ev1.map.str":     "str1",
		"ev1.map.int":     2,
		"ev1.map.map":     e.attrs["map"].(map[string]interface{})["map"],
		"ev1.map.map.str": "str2",
		"ev1.map.slice":   []string{"1", "2", "3"},
		"ev1.map.slice.0": "1",
		"ev1.map.slice.1": "2",
		"ev1.map.slice.2": "3",

		"ev1.nonexistant":             nil,
		"ev1.str.nonexistant":         nil,
		"ev1.map.str.nonexistant":     nil,
		"ev1.map.int.nonexistant":     nil,
		"ev1.map.nonexistant":         nil,
		"ev1.map.map.nonexistant":     nil,
		"ev1.map.map.str.nonexistant": nil,
		"ev1.map.slice.nonexistant":   nil,
		"ev1.map.slice.10":            nil,
	}

	for keyPath, expectedResult := range lookups {
		impl := attributeLookup(keyPath)
		result, err := impl.Value(CapturedEvents{
			"ev1": e,
		})

		if expectedResult == nil {
			assert.Nil(t, result, fmt.Sprintf("Unexpected result evaluating keyPath %s", keyPath))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error evaluating keyPath %s", keyPath))
			assert.Equal(t, expectedResult, result, fmt.Sprintf("Unexpected result evaluating keyPath %s", keyPath))
		}
	}
}
