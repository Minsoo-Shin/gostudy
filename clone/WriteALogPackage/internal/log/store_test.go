package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var (
	write = []byte("hello my name is minsoo111")
	width = uint64(len(write)) + lenWidth
)

func Test_OsStat(t *testing.T) {
	fi, err := os.Stat("../static/text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fi.Mode())
	fmt.Println(fi.IsDir())
	fmt.Println(fi.Size())
	fmt.Println(fi.Name())
	fmt.Println(fi.Sys())
	fmt.Println(fi.ModTime())
}

func Test_store_AppendRead(t *testing.T) {
	f, err := os.CreateTemp("", "append_read_test")
	assert.NoError(t, err)
	defer os.Remove(f.Name())

	s, err := newStore(f)
	assert.NoError(t, err)

	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	s, err = newStore(f)
	assert.NoError(t, err)
	testRead(t, s)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(write)
		assert.NoError(t, err)
		assert.Equal(t, pos+n, width*i)
	}
}

func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ {
		read, err := s.Read(pos)
		assert.NoError(t, err)
		assert.Equal(t, write, read)
		pos += width
	}
}

func testReadAt(t *testing.T, s *store) {
	t.Helper()
	for i, off := uint64(1), int64(0); i < 4; i++ {
		b := make([]byte, lenWidth)
		n, err := s.ReadAt(b, off)
		assert.NoError(t, err)
		assert.Equal(t, lenWidth, n)
		off += int64(n)

		size := enc.Uint64(b)
		b = make([]byte, size)
		n, err = s.ReadAt(b, off)
		assert.NoError(t, err)
		assert.Equal(t, write, b)
		assert.Equal(t, int(size), n)
		off += int64(n)
		fmt.Println(string(b))
	}
}
