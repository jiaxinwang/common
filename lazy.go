package common

import (
	"strings"
)

// Lazy ...
func Lazy(params map[string]interface{}) (eq, gt, lt, gteq, lteq map[string]interface{}) {
	eq = make(map[string]interface{})
	gt = make(map[string]interface{})
	lt = make(map[string]interface{})
	gteq = make(map[string]interface{})
	lteq = make(map[string]interface{})

	for k, v := range params {
		name := k
		dest := &eq
		switch {
		case strings.HasSuffix(k, `_gt`):
			name = strings.TrimSuffix(k, `_gt`)
			dest = &gt
		case strings.HasSuffix(k, `_lt`):
			name = strings.TrimSuffix(k, `_lt`)
			dest = &lt
		case strings.HasSuffix(k, `_gteq`):
			name = strings.TrimSuffix(k, `_gteq`)
			dest = &gteq
		case strings.HasSuffix(k, `_lteq`):
			name = strings.TrimSuffix(k, `_lteq`)
			dest = &lteq
		}
		(*dest)[name] = v
	}

	return
}
