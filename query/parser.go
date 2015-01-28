package query

import (
	"fmt"
	"strconv"
	"time"

	// log "github.com/cihub/seelog"
)

func visitToken(t *token) (*token, error) {
	newChildren := make([]*token, 0, len(t.children))
	skipNext := false

	for i, child := range t.children {
		if skipNext {
			skipNext = false
			continue
		}

		switch child.tt {
		case ttConjunction, ttDisjunction, ttEq, ttNe, ttGt, ttLt, ttGe, ttLe:
			if i > 0 && len(t.children) > i+1 {
				child.children = append(child.children, newChildren[len(newChildren)-1], t.children[i+1])
				newChildren = newChildren[:len(newChildren)-1]
				skipNext = true
			}
		}

		if newChild, err := visitToken(child); err != nil {
			return nil, err
		} else {
			newChildren = append(newChildren, newChild)
		}
	}
	t.children = newChildren
	return t, nil
}

func postprocessTokens(tokens []*token) ([]*token, error) {
	// This function is, honestly, cruft. Should be more efficient.
	whereToken := &token{
		tt:       ttWhereClause,
		children: make([]*token, 0, 2),
	}
	result := make([]*token, 0, len(tokens))

	for _, t := range tokens {
		switch t.tt {
		// Crufty: find all root-level predicates and connectives, stick them inside a ttWhereClause token.
		case ttPredicate, ttConjunction, ttDisjunction:
			whereToken.children = append(whereToken.children, t)

		default:
			result = append(result, t)
		}
	}

	if len(whereToken.children) > 0 {
		result = append(result, whereToken)
	}

	for i, t := range result {
		if newToken, err := visitToken(t); err != nil {
			return nil, err
		} else {
			result[i] = newToken
		}
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

func parseValue(t *token) (value, error) {
	switch t.tt {
	case ttAttributeSelector:
		return attributeLookup(t.content), nil

	case ttStringLiteral:
		return literalValue{t.content}, nil

	case ttNumericLiteral:
		if val, err := strconv.ParseFloat(t.content, 32); err != nil {
			return nil, err
		} else {
			return literalValue{val}, nil
		}

	default:
		return nil, fmt.Errorf("Unhandled token type: %s", t.tt.String())
	}
}

func parseWhereClauseToken(t *token) (Predicate, error) {
	parseChildren := func() ([]Predicate, error) {
		results := make([]Predicate, 0, len(t.children))
		for _, child := range t.children {
			if r, err := parseWhereClauseToken(child); err != nil {
				return nil, err
			} else if r != nil {
				results = append(results, r)
			}
		}
		if len(results) > 0 {
			return results, nil
		} else {
			return nil, nil
		}
	}

	switch t.tt {
	case ttEq, ttNe, ttGt, ttLt, ttGe, ttLe:
		if len(t.children) != 2 {
			return nil, fmt.Errorf("%s must have 2 children", t.tt.String())
		}

		if left, err := parseValue(t.children[0]); err != nil {
			return nil, err
		} else if right, err := parseValue(t.children[1]); err != nil {
			return nil, err
		} else {
			result := &operatorPredicate{
				left:  left,
				right: right,
			}

			switch t.tt {
			case ttEq:
				result.op = opEq
			case ttNe:
				result.op = opNe
			case ttGt:
				result.op = opGt
			case ttLt:
				result.op = opLt
			case ttGe:
				result.op = opGe
			case ttLe:
				result.op = opLe
			}

			return result, nil
		}

	case ttConjunction:
		if childCaptures, err := parseChildren(); err != nil {
			return nil, err
		} else if childCaptures == nil || len(childCaptures) < 1 {
			return nil, nil
		} else {
			return append(make(conjunction, 0, len(childCaptures)), childCaptures...), nil
		}

	case ttDisjunction:
		if childCaptures, err := parseChildren(); err != nil {
			return nil, err
		} else if childCaptures == nil || len(childCaptures) < 1 {
			return nil, nil
		} else {
			return append(make(disjunction, 0, len(childCaptures)), childCaptures...), nil
		}

	case ttPredicate, ttWhereClause:
		if result, err := parseChildren(); err != nil || result == nil || len(result) < 1 {
			return nil, err
		} else {
			return result[0], nil
		}

	default:
		return nil, fmt.Errorf("Unhandled token type: %s", t.tt.String())
	}
}

func parseWithinClauseToken(t *token) (time.Duration, error) {
	parseChildren := func() (time.Duration, error) {
		for _, child := range t.children {
			return parseWithinClauseToken(child)
		}
		return time.Duration(0), nil
	}

	switch t.tt {
	case ttDuration:
		return time.ParseDuration(t.content)

	case ttWithinClause:
		return parseChildren()

	default:
		return time.Duration(0), fmt.Errorf("Unhandled token type: %s", t.tt.String())
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
				q.capture = capture
			}

		case ttWhereClause:
			if predicate, err := parseWhereClauseToken(t); err != nil {
				return nil, fmt.Errorf("Error parsing %s: %s", t.tt.String(), err.Error())
			} else {
				q.predicate = predicate
			}

		case ttWithinClause:
			if window, err := parseWithinClauseToken(t); err != nil {
				return nil, fmt.Errorf("Error parsing %s: %s", t.tt.String(), err.Error())
			} else {
				q.window = window
			}

		default:
			return nil, fmt.Errorf("Unhandled root token type: %s", t.tt.String())
		}
	}

	if err := q.validate(); err != nil {
		return nil, err
	} else {
		return q, nil
	}
}

func Parse(data string) (*Query, error) {
	// start := time.Now()
	// defer log.Tracef("[sase:Parser] Took %s", time.Since(start).String())

	tokens, err := tokenize(data)
	if err != nil {
		return nil, fmt.Errorf("Error tokenizing: %s", err.Error())
	}
	tokens, err = postprocessTokens(tokens)
	if err != nil {
		return nil, fmt.Errorf("Error postprocessing tokens: %s", err.Error())
	}
	// for _, t := range tokens {
	// 	log.Tracef("[Parser] Parsed input tokens: %s", t.Tree())
	// }
	return parseTokens(tokens)
}
