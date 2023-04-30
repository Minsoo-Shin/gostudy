package main

import "testing"

func TestCheckSliceType(t *testing.T) {
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
				any: &Pizza{
					{
						Name: "빵",
					},
					{
						Name: "페페로니",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckSliceType(tt.args.any)
		})
	}
}
