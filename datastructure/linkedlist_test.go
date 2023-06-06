package datastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_PushBack(t *testing.T) {
	l := LinkedList[int]{}

	l.PushBack(100)
	assert.Equal(t, 100, l.root.value)

	l.PushBack(102)
	assert.Equal(t, 102, l.root.next.value)

	l.PushBack(104)
	assert.Equal(t, 104, l.root.next.next.value)
}

func TestLinkedList_PushFront(t *testing.T) {
	l := LinkedList[int]{}

	l.PushFront(100)
	assert.Equal(t, 100, l.root.value)

	l.PushFront(102)
	assert.Equal(t, 102, l.root.value)

	l.PushFront(104)
	assert.Equal(t, 104, l.root.value)
}

func TestLinkedList_First(t *testing.T) {
	l := LinkedList[int]{}

	l.PushFront(100)
	assert.Equal(t, 100, l.First().value)

	l.PushFront(101)
	assert.Equal(t, 101, l.First().value)
}

func TestLinkedList_Last(t *testing.T) {
	l := LinkedList[int]{}

	l.PushBack(100)
	assert.Equal(t, 100, l.Last().value)

	l.PushBack(101)
	assert.Equal(t, 101, l.Last().value)
}

func TestLinkedList_Count(t *testing.T) {
	l := LinkedList[int]{}

	l.PushFront(100)
	assert.Equal(t, 1, l.Count())

	l.PushBack(101)
	assert.Equal(t, 2, l.Count())
}

func TestLinkedList_GetAt(t *testing.T) {
	l := LinkedList[int]{}

	l.PushBack(100)
	l.PushBack(102)
	l.PushBack(104)

	assert.Equal(t, 100, l.GetAt(0).value)
	assert.Equal(t, 102, l.GetAt(1).value)
	assert.Equal(t, 104, l.GetAt(2).value)
}

func TestLinkedList_InsertAfter(t *testing.T) {
	l := LinkedList[int]{}

	l.PushBack(100)
	l.PushBack(102)
	l.PushBack(104)

	l.InsertAfter(l.GetAt(2), 200)
	assert.Equal(t, 200, l.GetAt(3).value)

	node := &Node[int]{
		value: 10,
	}

	l.InsertAfter(node, 300)
	assert.Equal(t, 4, l.Count())

}
