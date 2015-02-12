package query

import (
	"time"
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
