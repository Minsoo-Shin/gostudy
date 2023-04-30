package main

import (
	"fmt"
	"reflect"
)

// pizza equal toppings
type Pizza []Topping

type Topping struct {
	Name string
}

func CheckSliceType(any interface{}) {
	ve := reflect.ValueOf(any).Elem()
	fmt.Println(ve.Kind())
	for i := 0; i < ve.Len(); i++ {
		value := ve.Index(i)
		fmt.Printf("타입:%v, 값:%v\n", value.Type(), value.Interface())
	}
}
