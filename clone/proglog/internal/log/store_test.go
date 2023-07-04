package log

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	write = []byte("hello my name is minsoo111")
	width = uint64(len(write)) + lenWidth
)

const (
	TotalDataCount  = 100
	MaxStringLength = 100
	BufferSize      = 4096
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

func Test_Write(t *testing.T) {
	f, _ := os.CreateTemp("", "")
	buf := bufio.NewWriter(f)
	fmt.Println(buf.WriteString("why"))
	binary.Write(buf, enc, uint64(2))

	//n, err := buf.WriteString("a")
	fileInfo, _ := f.Stat()
	fmt.Println("a", fileInfo.Size())
	buf.Flush()
	fmt.Println("a", fileInfo.Size())
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

func TestStore_Close(t *testing.T) {
	f, err := os.CreateTemp("", "store_close_test")
	assert.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	assert.NoError(t, err)
	_, _, err = s.Append(write)
	assert.NoError(t, err)

	_, beforeSize, err := openFile(f.Name())
	assert.NoError(t, err)

	err = s.Close()
	assert.NoError(t, err)

	_, afterSize, err := openFile(f.Name())
	assert.NoError(t, err)
	assert.Equal(t, true, afterSize > beforeSize)
	t.Logf("before %d, afterSize %d", beforeSize, afterSize)
}

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fi.Size(), nil
}

func Test_fileAndBufferSpeed(t *testing.T) {
	// 파일에 직접 쓰는 방식의 시간 측정
	startTime := time.Now()
	writeDirect()
	directDuration := time.Since(startTime)

	// 버퍼를 사용하여 쓰는 방식의 시간 측정
	startTime = time.Now()
	writeBuffered()
	bufferedDuration := time.Since(startTime)

	fmt.Printf("파일에 직접 쓰기: %v\n", directDuration)
	fmt.Printf("버퍼를 사용하여 쓰기: %v\n", bufferedDuration)
}

// 파일에 직접 쓰는 방식
func writeDirect() {
	file, err := os.Create("direct_output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < TotalDataCount; i++ {
		data := generateRandomString() + "\n"
		_, err := file.WriteString(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// 버퍼를 사용하여 쓰는 방식
func writeBuffered() {
	file, err := os.Create("buffered_output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, BufferSize)

	for i := 0; i < TotalDataCount; i++ {
		data := generateRandomString() + "\n"
		_, err := writer.WriteString(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 버퍼가 가득 차면 Flush() 호출하여 데이터를 파일에 쓰기
		if writer.Buffered() >= BufferSize {
			err = writer.Flush()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// 마지막으로 남은 데이터를 파일에 쓰기
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(MaxStringLength) + 1
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(rand.Intn(26) + 97) // 소문자 알파벳
	}

	return string(b)
}
