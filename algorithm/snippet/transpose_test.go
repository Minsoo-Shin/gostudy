package snippet

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		name      string
		twoDArray [][]int
		want      [][]int
	}{
		{
			name: "square",
			twoDArray: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			name: "1x3 => 3x1",
			twoDArray: [][]int{
				{1, 2, 3},
			},
			want: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "3x2 => 2x3",
			twoDArray: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transpose(tt.twoDArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
