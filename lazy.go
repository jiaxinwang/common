package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// LazyParse ...
func LazyParse(v string, k reflect.Kind) (ret interface{}) {
	switch k {
	case reflect.Uint:
		if kv, err := strconv.ParseUint(v, 10, 64); err == nil {
			ret = uint(kv)
		}
	case reflect.Uint64:
		if kv, err := strconv.ParseUint(v, 10, 64); err == nil {
			ret = uint64(kv)
		}
	case reflect.Uint32:
		if kv, err := strconv.ParseUint(v, 10, 32); err == nil {
			ret = uint32(kv)
		}
	case reflect.Uint16:
		if kv, err := strconv.ParseUint(v, 10, 16); err == nil {
			ret = uint16(kv)
		}
	case reflect.Uint8:
		if kv, err := strconv.ParseUint(v, 10, 8); err == nil {
			ret = uint8(kv)
		}
	case reflect.Int:
		if kv, err := strconv.ParseInt(v, 10, 64); err == nil {
			ret = int(kv)
		}
	case reflect.Int64:
		if kv, err := strconv.ParseInt(v, 10, 64); err == nil {
			ret = int64(kv)
		}
	case reflect.Int32:
		if kv, err := strconv.ParseInt(v, 10, 32); err == nil {
			ret = int32(kv)
		}
	case reflect.Int16:
		if kv, err := strconv.ParseInt(v, 10, 16); err == nil {
			ret = int16(kv)
		}
	case reflect.Int8:
		if kv, err := strconv.ParseInt(v, 10, 8); err == nil {
			ret = int8(kv)
		}
	case reflect.String:
		ret = v
	default:
		fmt.Print("unsupported kind")
	}
	return
}

// LazyTagSlice ...
func LazyTagSlice(v interface{}, m map[string][]string) map[string][]interface{} {
	ret := make(map[string][]interface{})
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag
		if t := tag.Get(`lazy`); t != `` {
			if vv, ok := m[t]; ok {
				ret[t] = make([]interface{}, 0)
				for _, vvv := range vv {
					ret[t] = append(ret[t], LazyParse(vvv, field.Type.Kind()))
				}
			}
		}
	}

	return ret
}

// LazyTag ...
func LazyTag(v interface{}, m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag
		if t := tag.Get(`lazy`); t != `` {
			if v, ok := m[t]; ok {
				ret[t] = LazyParse(v, field.Type.Kind())
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
