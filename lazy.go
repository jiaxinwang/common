package common

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

// LazyStructMap ...
func LazyStructMap(src interface{}, timeLayout string) (ret map[string]interface{}, err error) {
	if b, err := json.Marshal(src); err != nil {
		return nil, err
	} else {
		ret = make(map[string]interface{})
		if err = json.Unmarshal(b, &ret); err != nil {
			return nil, err
		}

		switch v := reflect.ValueOf(src); v.Kind() {
		case reflect.Struct:
			tofs := reflect.TypeOf(src)
			vofs := reflect.ValueOf(src)
			for i := 0; i < vofs.NumField(); i++ {
				switch vofs.Field(i).Interface().(type) {
				case *time.Time:
					t := vofs.Field(i).Interface().(*time.Time)
					name := tofs.Field(i).Tag.Get(`lazy`)
					ret[name] = t.Format(timeLayout)
				case time.Time:
					t := vofs.Field(i).Interface().(time.Time)
					name := tofs.Field(i).Tag.Get(`lazy`)
					ret[name] = t.Format(timeLayout)
				}
			}
		default:
		}

		return ret, nil
	}
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}

// LazyMapStruct ...
func LazyMapStruct(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         nil,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

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
	case reflect.Bool:
		if kv, err := strconv.ParseBool(v); err == nil {
			ret = kv
		}
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
	logrus.Print(m)
	ret := make(map[string]interface{})
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag
		if t := tag.Get(`lazy`); t != `` {
			if vv, ok := m[t]; ok {
				name := field.Name
				r := reflect.ValueOf(v)
				f := reflect.Indirect(r).FieldByName(name)
				fieldValue := f.Interface()
				switch vvv := fieldValue.(type) {
				case uint64:
					i, _ := strconv.ParseUint(vv, 10, 64)
					ret[t] = i
				case uint32:
					i, _ := strconv.ParseUint(vv, 10, 32)
					ret[t] = i
				case uint:
					i, _ := strconv.ParseUint(vv, 10, 64)
					ret[t] = int(i)
				case int64:
					i, _ := strconv.ParseInt(vv, 10, 64)
					ret[t] = i
				case int32:
					i, _ := strconv.ParseInt(vv, 10, 32)
					ret[t] = i
				case int:
					i, _ := strconv.ParseInt(vv, 10, 64)
					ret[t] = int(i)
				case string:
					ret[t] = vv
				case bool:
					ret[t], _ = strconv.ParseBool(vv)
				case time.Time:
					ret[t], _ = time.Parse(time.RFC3339, vv)
				default:
					_ = vvv
				}
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

// LazyURLValues ...
func LazyURLValues(s interface{}, q url.Values) (eqm map[string][]interface{}, gtm, ltm, gtem, ltem map[string]interface{}) {
	eq, gt, lt, gte, lte := Lazy(q)
	eqm = LazyTagSlice(s, eq)
	gtm = LazyTag(s, gt)
	ltm = LazyTag(s, lt)
	gtem = LazyTag(s, gte)
	ltem = LazyTag(s, lte)
	return
}

// SelectBuilder ...
func SelectBuilder(s sq.SelectBuilder, eq map[string][]interface{}, gt, lt, gte, lte map[string]interface{}) sq.SelectBuilder {
	for k, v := range eq {
		switch {
		case len(v) == 1:
			eqs := sq.Eq{k: v[0]}
			s = s.Where(eqs)
		case len(v) > 1:
			eqs := sq.Eq{k: v}
			s = s.Where(eqs)
		}
	}
	if len(gt) > 0 {
		m := sq.Gt(gt)
		s = s.Where(m)
	}
	if len(lt) > 0 {
		m := sq.Lt(lt)
		s = s.Where(m)
	}
	if len(gte) > 0 {
		m := sq.GtOrEq(gte)
		s = s.Where(m)
	}
	if len(lte) > 0 {
		m := sq.LtOrEq(lte)
		s = s.Where(m)
	}
	return s
}

// Query ...
func Query(db *gorm.DB, active sq.SelectBuilder) (ret []map[string]interface{}, err error) {
	ret = make([]map[string]interface{}, 0)
	sql, args, err := active.ToSql()
	if err != nil {
		return ret, err
	}

	rows, sqlErr := db.Raw(sql, args...).Rows()

	defer rows.Close()
	if sqlErr != nil {
		return ret, sqlErr
	}

	columns, err := rows.Columns()
	if err != nil {
		return ret, err
	}
	length := len(columns)
	for rows.Next() {
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			return ret, err
		}
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			k := columns[i]
			val := *(current[i]).(*interface{})
			if val == nil {
				value[k] = nil
				continue
			}
			vType := reflect.TypeOf(val)
			switch vType.String() {
			case "uint8":
				value[k] = val.(int8)
			case "uint16":
				value[k] = val.(int16)
			case "uint32":
				value[k] = val.(int32)
			case "uint64":
				value[k] = val.(int64)
			case "int8":
				value[k] = val.(int8)
			case "int16":
				value[k] = val.(int16)
			case "int32":
				value[k] = val.(int32)
			case "int64":
				value[k] = val.(int64)
			case "bool":
				value[k] = val.(bool)
			case "string":
				value[k] = val.(string)
			case "time.Time":
				value[k] = val.(time.Time)
			case "[]uint8":
				value[k] = string(val.([]uint8))
			default:
				// logrus.Warnf("unsupport data type '%s' now\n", vType)
			}
		}
		ret = append(ret, value)
	}

	return ret, nil

}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
