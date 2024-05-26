package util

import (
	"reflect"
)

func InitEmptySlice(any interface{}) {
	target := reflect.ValueOf(any)
	elems := target.Elem()
	switch elems.Kind() {
	case reflect.Slice:
		if elems.IsNil() {
			elems.Set(reflect.MakeSlice(elems.Type(), 0, 0))
			return
		}

		for i := 0; i < elems.Len(); i++ {
			InitEmptySlice(elems.Index(i).Addr().Interface())
		}

	case reflect.Struct:
		for i := 0; i < elems.NumField(); i++ {
			fieldValue := elems.Field(i)
			InitEmptySlice(fieldValue.Addr().Interface())
		}
	}
}
