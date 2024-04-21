package snippet

type LinkedList struct {
	root *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		root: nil,
		tail: nil,
	}
}

func (l *LinkedList) PushBack(val int) *LinkedList {
	node := &Node{
		val:  val,
		next: nil,
	}
	if l.root == nil {
		l.root = node
		l.tail = node
		return l
	}
	// 현재 tail의 next에 추가되는 node 주소 저장
	l.tail.next = node
	// tail은 추가되는 node의 주소로 변경
	l.tail = node
	return l
}

func (l *LinkedList) PushFront(val int) *LinkedList {
	node := &Node{
		val:  val,
		next: nil,
	}
	if l.root == nil {
		l.root = node
		l.tail = node
		return l
	}
	// 추가되는 node.next가 현재 root의 node
	node.next = l.root
	// l.root는 새롭게 추가된 node가 된다.
	l.root = node
	return l
}

func (l *LinkedList) Front() *Node {
	return l.root
}

func (l *LinkedList) Back() *Node {
	return l.tail
}

func (l *LinkedList) Count() int {
	cur := l.root
	cnt := 0

	for ; cur != nil; cur = cur.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList) GetAt(idx int) *Node {
	cur := l.root
	curIdx := 0
	for ; curIdx < idx; curIdx++ {
		cur = cur.next
	}
	return cur
}
