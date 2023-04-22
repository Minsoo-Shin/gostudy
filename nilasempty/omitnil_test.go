package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type ListPersons struct {
	Persons []Person
}

type Person struct {
	Name          string
	Age           string
	ProfileImages []Image
}

type Image struct {
	URL  string
	Tags []Tag
}

type Tag struct {
	AnySlice []string
}

func TestInitEmptySlice(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				any: &ListPersons{},
			},
		},
		{
			name: "test2",
			args: args{
				any: &Person{
					Name:          "minsoo",
					Age:           "30",
					ProfileImages: nil,
				},
			},
		},
		{
			name: "test3",
			args: args{
				any: &Person{
					Name: "minsoo",
					Age:  "30",
					ProfileImages: []Image{
						{
							URL:  "www.test.com/아무나/코딩",
							Tags: nil,
						},
					},
				},
			},
		},
		{
			name: "test4",
			args: args{
				any: &Person{
					Name: "minsoo",
					Age:  "30",
					ProfileImages: []Image{
						{
							URL: "www.test.com/아무나/코딩",
							Tags: []Tag{
								{
									AnySlice: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintJSON(tt.args.any)
			InitEmptySlice(tt.args.any)
			PrintJSON(tt.args.any)
		})
	}
}

func PrintJSON(any interface{}) {
	data, _ := json.Marshal(any)
	fmt.Println(string(data))
}

func Test_CheckReflect(t *testing.T) {
	fmt.Println(reflect.TypeOf(12))
	fmt.Println(reflect.TypeOf("12"))
	fmt.Println(reflect.TypeOf([]string{"12", "23"}))
	fmt.Println(reflect.TypeOf(map[string]string{}))

	fmt.Println(reflect.ValueOf(12))
	fmt.Println(reflect.ValueOf("12"))
	fmt.Println(reflect.ValueOf([]string{"12", "23"}))
	fmt.Println(reflect.ValueOf(map[string]string{}))
}
