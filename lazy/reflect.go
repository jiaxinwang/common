package lazy

import (
	"reflect"
	"strings"
)

func valueOfTag(inter interface{}, tagName string) interface{} {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if strings.EqualFold(t.Field(i).Tag.Get("lazy"), tagName) {
				return v.Field(i).Interface()
			}
		}
	}
	return nil
}

func foreignOfModel(inter interface{}) [][3]string {
	ret := make([][3]string, 0)
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	val := reflect.Indirect(reflect.ValueOf(inter))
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			tag := t.Field(i).Tag.Get("lazy")
			if len(tag) > 0 {
				if _, table, id, err := disassembleTag(tag); err == nil && len(table) != 0 && len(id) != 0 {
					f := [3]string{val.Type().Field(i).Name, table, id}
					ret = append(ret, f)
				}
			}
		}
	}
	return ret
}
