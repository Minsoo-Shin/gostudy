package main

import "testing"

type personA struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func TestMetadata(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				any: &personA{
					Name:  "minsoo",
					Age:   30,
					Hobby: []string{"nil없애기", "심심코딩"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Metadata(tt.args.any)
		})
	}
}
