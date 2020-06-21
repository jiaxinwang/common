package common

import (
	"strings"
)

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
