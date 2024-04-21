package snippet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateLeftByOne1DArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "원소 갯수가 5인 배열",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{2, 3, 4, 5, 1},
		},
		{
			name:  "원소 갯수가 0인 배열",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateLeftByOne1DArray(tt.input); !reflect.DeepEqual(got, tt.want) {
				fmt.Errorf("got: %v, want: %v", got, tt.want)
			}

		})
	}
}

func Test_rotateRightByOne1DArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "길이가 5인 slice",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRightByOne1DArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRightByOne1DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateLeftByCount(t *testing.T) {
	type args struct {
		arr   []int
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "원소가 5개이고, 3개 left rotate",
			args: args{
				arr:   []int{1, 2, 3, 4, 5},
				count: 3,
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "원소가 5개이고, 6개 left rotate",
			args: args{
				arr:   []int{1, 2, 3, 4, 5},
				count: 6,
			},
			want: []int{2, 3, 4, 5, 1},
		},
		{
			name: "원소가 0개인 경우",
			args: args{
				arr:   []int{},
				count: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateLeftByCount(tt.args.arr, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateLeftByCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
