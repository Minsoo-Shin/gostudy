package datastructure

/*

단반향 연결 리스트(single linked list)를 구현해보자.

단방향 연결 리스트란, 데이터를 가진 노드들이 한 방향으로 연결되어있는 자료구조이다.
노드는 데이터값(value)과 다음 노드의 주소(pointer)를 갖고 있다.
그래서 노드는 한 방향으로 탐색이 가능하여 단반향 연결 리스트라고 한다.

type Node[T any] struct {
	Value T
	Next *Node[T]
}
*/

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func New() LinkedList[int] {
	return LinkedList[int]{
		root:  nil,
		tail:  nil,
		count: 0,
	}
}

func (l *LinkedList[T]) PushBack(value T) {
	back := &Node[T]{
		value: value,
	}

	l.count++
	if l.root == nil {
		l.root = back
		l.tail = back
		return
	}

	l.tail.next = back
	l.tail = back
}

func (l *LinkedList[T]) PushFront(value T) {
	front := &Node[T]{
		value: value,
	}
	l.count++
	if l.root == nil {
		l.root = front
		l.tail = front
		return
	}

	front.next = l.root
	l.root = front
}

func (l *LinkedList[T]) First() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Last() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(index int) *Node[T] {
	if index >= l.Count() {
		return nil
	}

	var i int
	for node := l.root; node != nil; node = node.next {
		if i == index {
			return node
		}
		i++
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.includeNode(node) {
		return
	}
	newNode := &Node[T]{
		value: value,
		next:  nil,
	}

	node.next, newNode.next = newNode, node.next
	l.count++
}

func (l *LinkedList[T]) includeNode(node *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == node {
			return true
		}
	}

	return false
}

// [node1] - [node2] - [node3] , [newNode]R
func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	// node가 root면 찾을 수 없다.
	if l.root == node {
		l.PushFront(value)
	}

	newNode := &Node[T]{
		value: value,
		next:  nil,
	}
	// node 앞을 찾아야한다.
	prevNode := l.findPrevNode(node)
	prevNode.next, newNode.next = newNode, prevNode.next
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	panic("implement")
}

func (l *LinkedList[T]) PopFront() {
	panic("implement")
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	panic("implement")
}
