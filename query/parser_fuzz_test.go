package query

import (
	"bytes"
	"crypto/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func fuzzy(length int) []byte {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

func TestParsingComponentFuzz(t *testing.T) {
	cases := []string{
		"EVENT xxx xxx",
		"EVENT SEQ(xxx xxx, xxx xxx)",
		"EVENT ANY(xxx xxx, xxx xxx, xxx xxx)",
		"EVENT a b WITHIN xxxh",
		"EVENT xxx xxx WITHIN xxxh",
		"EVENT a b WHERE a.b = \"xxx\"",
	}

	for _, unformatted := range cases {
		for i := 0; i < 100; i++ {
			buf := new(bytes.Buffer)
			for ii, part := range strings.Split(unformatted, "xxx") {
				if ii > 1 {
					buf.Write(fuzzy(100))
				}
				buf.WriteString(part)
			}

			require.NotPanics(t, func() {
				Parse(buf.String())
			})
		}
	}
}
