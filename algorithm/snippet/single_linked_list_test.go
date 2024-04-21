package snippet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Count(t *testing.T) {

	tests := []struct {
		name string
		ll   *LinkedList
		want int
	}{
		{
			name: "빈 링크드 리스트",
			ll:   NewLinkedList(),
			want: 0,
		},
		{
			name: "링크드 리스트",
			ll:   NewLinkedList().PushBack(1),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Count(); got != tt.want {
				t.Errorf("LinkedList.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_GetAt(t *testing.T) {
	ll := NewLinkedList()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	got := ll.Count()
	assert.Equal(t, 3, got)

	assert.Equal(t, 1, ll.GetAt(0).val)
	assert.Equal(t, 2, ll.GetAt(1).val)
	assert.Equal(t, 3, ll.GetAt(2).val)
}
