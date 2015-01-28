package query

import (
	"fmt"
	"testing"

	log "github.com/cihub/seelog"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	// Don't put semicolons on the end; that happens automatically
	expectations := map[string]bool{
		// EVENT-only
		"EVENT a b":                                 true,
		"EVENT SEQ(a b)":                            true,
		"EVENT SEQ(a b, a d)":                       true,
		"EVENT ANY(a b)":                            true,
		"EVENT ANY(a b, c d)":                       true,
		"EVENT SEQ(a b, ANY(c d, e f))":             true,
		"EVENT SEQ(a e1, !(c e2), ANY(c e3, d e4))": true,
		// Errors
		"EVENT":               false, // No capture
		"EVENT a":             false, // No alias
		"EVENT 1a b":          false, // Identifiers must begin with alpha
		"EVENT a b, c d":      false, // No SEQ
		"EVENT SEQ()":         false, // Empty SEQ (no capture)
		"EVENT ANY()":         false, // Empty ANY (no capture)
		"EVENT SEQ(a b, c b)": false, // Clashing capture aliases

		// EVENT + WITHIN
		"EVENT a b WITHIN 1h":                                  true,
		"EVENT SEQ(a b) WITHIN 30m":                            true,
		"EVENT SEQ(a b, a d) WITHIN 2h30m20s":                  true,
		"EVENT ANY(a b) WITHIN 100h":                           true,
		"EVENT ANY(a b, c d) WITHIN 100000h":                   true,
		"EVENT SEQ(a b, ANY(c d, e f)) WITHIN 200h30m20s100ns": true,
		"EVENT SEQ(a e1, !(c e2), ANY(c e3, d e4)) WITHIN 1h":  true,
	}

	te := func(queryText string, expectSuccess bool) {
		assert.NotPanics(t, func() {
			q, err := Parse(queryText)
			if expectSuccess {
				assert.NoError(t, err, fmt.Sprintf("Unexpected error parsing \"%s\"", queryText))
				assert.NotNil(t, q, "Query unexpectedly nil for \"%s\"", queryText)

				// If we output the query again, and re-parse it, outputs should be the same
				// We can't compare strings directly because we deliberately standardise output
				output := q.QueryText()
				q2, err := Parse(output)
				assert.NoError(t, err, fmt.Sprintf("Unexpected error parsing generated output \"%s\" (original: \"%s\")",
					output, queryText))
				assert.Equal(t, output, q2.QueryText(), fmt.Sprintf("Generated outputs do not match for input \"%s\"",
					queryText))
			} else {
				assert.Error(t, err, fmt.Sprintf("Error expected parsing \"%s\"", queryText))
				assert.Nil(t, q, "Query unexpectedly not-nil for \"%s\"", queryText)
			}
		}, fmt.Sprintf("Unexpected panic parsing \"%s\"", queryText))
	}

	for queryText, expectSuccess := range expectations {
		log.Tracef("[sase:TestParsing] Trying \"%s\"…", queryText)
		te(queryText, expectSuccess)
		te(queryText+";", expectSuccess)
		te(queryText+"     ;", expectSuccess)
	}
}

func BenchmarkParsing(b *testing.B) {
	queryText := "EVENT SEQ(t1 e1, ANY(t2 e2, t3 e3), !(t4 e4), t5, e5) WHERE e1.foo == e2.bar AND e3.baz == e4.boop WITHIN 2h;"
	for i := 0; i < b.N; i++ {
		Parse(queryText)
	}
}

func BenchmarkParsingParallel(b *testing.B) {
	queryText := "EVENT SEQ(t1 e1, ANY(t2 e2, t3 e3), !(t4 e4), t5, e5) WHERE e1.foo == e2.bar AND e3.baz == e4.boop WITHIN 2h;"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Parse(queryText)
		}
	})
}
