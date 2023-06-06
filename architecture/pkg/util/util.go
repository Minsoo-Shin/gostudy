package util

import (
	"fmt"
	"reflect"
	"time"
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

func GetTestDateKey(date int64) string {
	t := time.Unix(date, 0).Local()
	year, month, day := t.Date()

	return fmt.Sprintf("%d%02d%02d", year, month, day)
}
