package main

import (
	"fmt"
	"reflect"
)

func Metadata(any interface{}) {
	target := reflect.ValueOf(any)
	elements := target.Elem()

	fmt.Printf("Type: %s\n", target.Type()) // 구조체 타입명

	for i := 0; i < elements.NumField(); i++ {
		mValue := elements.Field(i)
		mType := elements.Type().Field(i)
		tag := mType.Tag

		fmt.Printf("%10s %10s ==> %10v, json: %10s\n",
			mType.Name,         // 이름
			mType.Type,         // 타입
			mValue.Interface(), // 값
			tag.Get("json"))    // json 태그
	}
}
