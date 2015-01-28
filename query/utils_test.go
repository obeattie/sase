package query

type tEventImpl struct {
	typ   string
	attrs map[string]interface{}
}

func (e *tEventImpl) Type() string {
	return e.typ
}

func (e *tEventImpl) Attributes() map[string]interface{} {
	return e.attrs
}
