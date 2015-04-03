package query

import (
	"fmt"
	"strings"
	"time"

	"github.com/obeattie/sase/domain"
)

type tEventImpl struct {
	typ   string
	attrs map[string]interface{}
	ts    time.Time
}

func (e *tEventImpl) Type() string {
	return e.typ
}

func (e *tEventImpl) Attributes() map[string]interface{} {
	return e.attrs
}

func (e *tEventImpl) When() time.Time {
	return e.ts
}

func describeStack(evs domain.CapturedEvents) string {
	results := make([]string, 0, len(evs))

	for alias, event := range evs {
		desc := fmt.Sprintf("%s:%s", alias, event.Type())
		results = append(results, desc)
	}

	return strings.Join(results, " ")
}
