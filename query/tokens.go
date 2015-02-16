package query

import (
	"bytes"
)

//go:generate stringer -type=tt

// tt = token type
type tt uint16

const (
	ttGeneric tt = iota

	ttNumericLiteral
	ttStringLiteral

	ttEventClause
	ttEventDecl
	ttEventDeclType
	ttEventDeclAlias
	ttSeqDecl
	ttNegatedDecl
	ttAnyDecl

	ttWhereClause
	ttEq
	ttNe
	ttGt
	ttLt
	ttGe
	ttLe
	ttAttributeSelector
	ttEquivalenceTest
	ttConjunction
	ttDisjunction
	ttPredicate

	ttWithinClause
	ttDuration
)

type token struct {
	tt       tt
	content  string
	children []*token
}

func (t *token) AddChild(c *token) {
	if t.children == nil {
		t.children = make([]*token, 0, 5)
	}
	t.children = append(t.children, c)
}

func (t *token) writeDesc(buf *bytes.Buffer, depth int) {
	for i := 0; i < depth; i++ {
		buf.WriteString("    ")
	}
	buf.WriteString(t.tt.String())
	if t.content != "" {
		buf.WriteString(": \"")
		buf.WriteString(t.content)
		buf.WriteRune('"')
	}
	buf.WriteRune('\n')
	for _, child := range t.children {
		child.writeDesc(buf, depth+1)
	}
}

func (t *token) Tree() string {
	buf := new(bytes.Buffer)
	t.writeDesc(buf, 0)
	return buf.String()
}

type proposedToken struct {
	*token
	i  int // length of the token stack, at proposal time
	pi int // length of the proposal stack, at proposal time
}
