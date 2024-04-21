package snippet

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		input func() MinStack
		cmd   func(MinStack) int
		want  int
	}{
		{
			input: func() MinStack {
				stack := Constructor()
				stack.Push(1)
				stack.Push(100)
				stack.Push(5)
				return stack
			},
			cmd: func(s MinStack) int {
				return s.GetMin()
			},
			want: 1,
		},
		{
			input: func() MinStack {
				stack := Constructor()
				stack.Push(5)
				stack.Push(100)
				stack.Push(10)
				stack.Push(1)
				stack.Pop()
				return stack
			},
			cmd: func(s MinStack) int {
				return s.GetMin()
			},
			want: 5,
		},
		{
			input: func() MinStack {
				stack := Constructor()
				stack.Push(5)
				stack.Push(1)
				return stack
			},
			cmd: func(s MinStack) int {
				return s.Top()
			},
			want: 1,
		},
		{
			input: func() MinStack {
				stack := Constructor()
				return stack
			},
			cmd: func(s MinStack) int {
				return s.Top()
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := tt.cmd(tt.input())
			assert.Equalf(t, tt.want, got, "MinStack() = %v, want %v", got, tt.want)
		})
	}
}
