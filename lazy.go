package common

import (
	"reflect"
	"strings"
)

// LazyTag ...
func LazyTag(v interface{}, m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	val := reflect.ValueOf(v).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag
		if t := tag.Get(`lazy`); t != `` {
			if v, ok := m[field.Name]; ok {
				ret[t] = v
			}
		}
	}
	return ret
}

// Lazy ...
func Lazy(params map[string]interface{}) (eq, gt, lt, gte, lte map[string]interface{}) {
	eq = make(map[string]interface{})
	gt = make(map[string]interface{})
	lt = make(map[string]interface{})
	gte = make(map[string]interface{})
	lte = make(map[string]interface{})

	for k, v := range params {
		if v == nil {
			continue
		}
		name := k
		dest := &eq
		switch {
		case strings.EqualFold(k, "size"):
			fallthrough
		case strings.EqualFold(k, "offset"):
			fallthrough
		case strings.EqualFold(k, "page"):
			dest = nil
		case strings.HasSuffix(k, `_gt`):
			name = strings.TrimSuffix(k, `_gt`)
			dest = &gt
		case strings.HasSuffix(k, `_lt`):
			name = strings.TrimSuffix(k, `_lt`)
			dest = &lt
		case strings.HasSuffix(k, `_gte`):
			name = strings.TrimSuffix(k, `_gte`)
			dest = &gte
		case strings.HasSuffix(k, `_lte`):
			name = strings.TrimSuffix(k, `_lte`)
			dest = &lte
		}
		if dest != nil {
			(*dest)[name] = v
		}
	}
	return
}
