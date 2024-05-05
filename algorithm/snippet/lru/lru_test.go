package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	c := Constructor(3)

	c.Put(1, 10)
	c.Put(2, 20)
	c.Put(3, 30)
	assert.Equal(t, 10, c.Get(1))
	assert.Equal(t, 20, c.Get(2))
	assert.Equal(t, 30, c.Get(3))

	// 넘치면 최근에 사용하지 않은 애들부터 날림
	c.Put(4, 40)
	assert.Equal(t, -1, c.Get(1))

	// 기존의 있는 값 변경하기
	c.Put(3, 300)
	assert.Equal(t, 300, c.Get(3))

}
