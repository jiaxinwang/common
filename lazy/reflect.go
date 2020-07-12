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
