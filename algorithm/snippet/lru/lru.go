package lru

import (
	"container/list"
)

type LRUCache struct {
	Queue    *list.List    // 사용한 순으로 정렬
	Items    map[int]*Node // store
	Capacity int           // 사이즈
}

type Node struct {
	Data   int
	KeyPtr *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Queue:    list.New().Init(),
		Items:    make(map[int]*Node, 0),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// value를 찾고, 있다면 Queue에 순서를 앞으로 옮기기
	// key가 없으면 -1 을 반환한다.
	item, ok := this.Items[key]
	if !ok {
		return -1
	}
	this.Queue.MoveToFront(item.KeyPtr)
	return item.Data
}

func (this *LRUCache) Put(key int, value int) {
	if item, ok := this.Items[key]; !ok {
		if this.Queue.Len() == this.Capacity {
			// 오래된 키를 삭제
			delete(this.Items, this.Queue.Back().Value.(int))
			this.Queue.Remove(this.Queue.Back())
		}
		this.Items[key] = &Node{
			Data:   value,
			KeyPtr: this.Queue.PushFront(key),
		}
	} else {
		// 데이터를 변경하고 맨 앞으로 수정
		item.Data = value
		this.Queue.MoveToFront(item.KeyPtr)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
