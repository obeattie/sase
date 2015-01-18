package parser

type tokenType uint

const (
	ttRoot tokenType = iota
	ttText
	ttStringLiteral
	ttDuration
)

type token struct {
	Tt          tokenType
	Value       string
	FirstChild  *token
	NextSibling *token
}
