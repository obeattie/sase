package query

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/obeattie/sase/domain"
)

var ErrEventNotFound = errors.New("Cannot find event")

type value interface {
	Representable
	Value(domain.CapturedEvents) (interface{}, error)
	usedAliases() []string
}

// A literalValue is a simple value that always returns a constant
// NOTE: This only, at present, supports strings and float64's
type literalValue struct {
	v interface{}
}

func (v literalValue) QueryText() string {
	switch val := v.v.(type) {
	case string:
		return fmt.Sprintf("%q", val)
	case float64:
		return fmt.Sprintf("%f", val)
	default:
		return ""
	}
}

func (v literalValue) Value(evs domain.CapturedEvents) (interface{}, error) {
	return v.v, nil
}

func (v literalValue) usedAliases() []string {
	return nil
}

// An attributeLookup represents the lookup of an attribute from an event (it handles nested key paths)
type attributeLookup string

func (p attributeLookup) QueryText() string {
	return string(p)
}

func (p attributeLookup) Value(evs domain.CapturedEvents) (interface{}, error) {
	parts := strings.Split(string(p), ".")

	if len(parts) > 1 {
		var result interface{}

		if ev, ok := evs[parts[0]]; !ok {
			return nil, ErrEventNotFound
		} else {
			result = ev.Attributes()
		}

		for _, part := range parts[1:] {
			val := reflect.ValueOf(result)
			typ := val.Type()

			switch typ.Kind() {
			case reflect.Map:
				reflectResult := val.MapIndex(reflect.ValueOf(part))
				if !reflectResult.IsValid() {
					return nil, fmt.Errorf("Attribute lookup failed for %s: cannot find map field %s", p, part)
				}
				result = reflectResult.Interface()

			case reflect.Array, reflect.Slice:
				if idx, err := strconv.Atoi(part); err != nil {
					return nil, fmt.Errorf("Attribute lookup failed for %s: %s", p, err.Error())
				} else if idx >= val.Len() {
					return nil, fmt.Errorf("Attribute lookup failed for %s: %d is beyond array/slice bounds", p,
						idx)
				} else {
					result = val.Index(idx).Interface()
				}

			default:
				return nil, fmt.Errorf("Attribute lookup failed for %s: cannot index type %s", p, typ.String())
			}
		}

		return result, nil
	} else {
		return nil, fmt.Errorf("Attribute lookup failed for %s: lookup has too few parts", p)
	}
}

func (p attributeLookup) usedAliases() []string {
	parts := strings.SplitN(string(p), ".", 2)
	if len(parts) < 1 {
		return nil
	}
	return parts[:1]
}
