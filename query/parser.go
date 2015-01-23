package query

import (
	"fmt"
)

func postprocessTokens(tokens []*token) ([]*token, error) {
	whereToken := &token{
		tt:       ttWhereClause,
		children: make([]*token, 0, 2),
	}
	result := make([]*token, 0, len(tokens))
	for _, t := range tokens {
		// Crufty: find all root-level predicates and connectives, stick them inside a ttWhereClause token.
		switch t.tt {
		case ttPredicate, ttConjunction, ttDisjunction:
			whereToken.children = append(whereToken.children, t)

		default:
			result = append(result, t)
		}
	}

	if len(whereToken.children) > 0 {
		result = append(result, whereToken)
	}

	return result, nil
}

func parseEventClauseToken(t *token) (EventCapture, error) {
	parseChildren := func() ([]EventCapture, error) {
		results := make([]EventCapture, len(t.children))
		for i, child := range t.children {
			if result, err := parseEventClauseToken(child); err != nil {
				return nil, err
			} else {
				results[i] = result
			}
		}
		return results, nil
	}

	switch t.tt {
	case ttEventClause:
		if len(t.children) > 0 {
			return parseEventClauseToken(t.children[0])
		} else {
			return nil, nil
		}

	case ttEventDecl:
		// MUST have a type and an alias, in that order
		return &basicEventCapture{
			eventType: t.children[0].content,
			name:      t.children[1].content,
		}, nil

	case ttSeqDecl:
		if childCaptures, err := parseChildren(); err != nil {
			return nil, err
		} else {
			return seqEventCapture(childCaptures), nil
		}

	case ttAnyDecl:
		if childCaptures, err := parseChildren(); err != nil {
			return nil, err
		} else {
			return anyEventCapture(childCaptures), nil
		}

	case ttNegatedDecl:
		if childCaptures, err := parseChildren(); err != nil {
			return nil, err
		} else if len(childCaptures) != 1 {
			return nil, fmt.Errorf("Negated declaration may only have one child")
		} else {
			return &negatedEventCapture{childCaptures[0]}, nil
		}

	default:
		return nil, fmt.Errorf("Unhandled token type: %s", t.tt.String())
	}
}

func parseTokens(tokens []*token) (*Query, error) {
	q := new(Query)

	for _, t := range tokens {
		switch t.tt {
		case ttEventClause:
			if capture, err := parseEventClauseToken(t); err != nil {
				return nil, fmt.Errorf("Error parsing %s: %s", t.tt.String(), err.Error())
			} else {
				q.Capture = capture
			}

		case ttWhereClause:

		case ttWithinClause:

		default:
			return nil, fmt.Errorf("Unhandled root token type: %s", t.tt.String())
		}
	}

	return q, nil
}

func Parse(data string) (*Query, error) {
	tokens, err := tokenize(data)
	if err != nil {
		return nil, fmt.Errorf("Error tokenizing: %s", err.Error())
	}
	tokens, err = postprocessTokens(tokens)
	if err != nil {
		return nil, fmt.Errorf("Error postprocessing tokens: %s", err.Error())
	}
	return parseTokens(tokens)
}
