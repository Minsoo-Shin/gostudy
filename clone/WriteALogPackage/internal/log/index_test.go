package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIndex(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "index_test")
	assert.NoError(t, err)
	defer os.Remove(f.Name())

	c := Config{}
	c.Segment.MaxIndexBytes = 1024
	idx, err := newIndex(f, c)
	assert.NoError(t, err)
	_, _, err = idx.Read(-1)
	assert.NoError(t, err)
	assert.Equal(t, f.Name(), idx.Name())
	entries := []struct {
		Off uint32
		Pos uint64
	}{
		{
			Off: 0,
			Pos: 0,
		},
		{
			Off: 1,
			Pos: 10,
		},
	}
	for _, want := range entries {
		err = idx.Write(want.Off, want.Pos)
		assert.NoError(t, err)

		_, pos, err := idx.Read(int64(want.Off))
		assert.NoError(t, err)
		assert.Equal(t, want.Pos, pos)

	}
}

// 파일 사이즈를 변환해주는 함수
func Test_Truncate(t *testing.T) {
	f, _ := os.Create("text.txt")

	fi, _ := f.Stat()
	fmt.Println(fi.Size())

	fmt.Println(os.Truncate(f.Name(), 2))

	fi, _ = f.Stat()
	fmt.Println(fi.Size())
}

func Test_FileDescriptor(t *testing.T) {
	/*
		파일 디스크립터는 컴퓨터 시스템에서 파일이나 입출력 장치를 식별하는 데 사용되는 숫자입니다. 파일은 디스크립터는 파일을 열거나 생성할 때, 생성되고, 프로세스에서 파일에 대한 작업을 수행하는데 사용한다고 함.
		실제로 파일을 생성해보고 파일 디스크립터를 프린트 해보자.

		default로 0~2까지는 표준 입력(stdin), 표준 출력(stdout), 표준 오류(stderr)를 나타낸다.
		그래서 아래 파일이 생성될 때, 3에서 부터 시작한다.
	*/

	f1, _ := os.CreateTemp("", "file3")
	defer f1.Close()
	fmt.Println(f1.Fd())

	f2, _ := os.CreateTemp("", "file4")
	defer f2.Close()
	fmt.Println(f2.Fd())

}
