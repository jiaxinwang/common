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
func Lazy(params map[string][]string) (eq map[string][]string, gt, lt, gte, lte map[string]string) {
	eq = make(map[string][]string)
	gt = make(map[string]string)
	lt = make(map[string]string)
	gte = make(map[string]string)
	lte = make(map[string]string)

	for kk, vv := range params {
		if vv == nil {
			continue
		}
		for _, v := range vv {
			name := kk
			destM := &eq
			destS := &gt
			switch {
			case strings.EqualFold(kk, "size"):
				fallthrough
			case strings.EqualFold(kk, "offset"):
				fallthrough
			case strings.EqualFold(kk, "page"):
				destM = nil
				destS = nil
			case strings.HasSuffix(kk, `_gt`):
				name = strings.TrimSuffix(kk, `_gt`)
				destM = nil
				destS = &gt
			case strings.HasSuffix(kk, `_lt`):
				name = strings.TrimSuffix(kk, `_lt`)
				destM = nil
				destS = &lt
			case strings.HasSuffix(kk, `_gte`):
				name = strings.TrimSuffix(kk, `_gte`)
				destM = nil
				destS = &gte
			case strings.HasSuffix(kk, `_lte`):
				name = strings.TrimSuffix(kk, `_lte`)
				destM = nil
				destS = &lte
			default:
				destM = &eq
				destS = nil
			}
			if destS != nil {
				(*destS)[name] = v
			}
			if destM != nil {
				if (*destM)[name] == nil {
					(*destM)[name] = make([]string, 0)
				}
				(*destM)[name] = append((*destM)[name], v)
			}
		}
	}
	return
}
