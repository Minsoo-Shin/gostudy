package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	api "gostudy/clone/StructureDataWithProtobuf/api/v1"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func Test_Ext(t *testing.T) {
	file, err := os.Create("text.txt")
	assert.NoErrorf(t, err, "%v", err)
	fmt.Println(path.Ext(file.Name()))

	offStr := strings.TrimSuffix(file.Name(), path.Ext(file.Name()))
	fmt.Println(offStr)

	off, _ := strconv.ParseUint(offStr, 10, 0)
	fmt.Println(off)
}

func TestLog(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T, log *Log,
	){
		"append and read a record succeeds": testAppendRead,
		"offset out of range error":         testOutOfRangeError,
		"init with existing segments":       testInitExisting,
		"reader":                            testReader,
		"truncate":                          testTruncate,
	} {
		t.Run(scenario, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "store-test")
			assert.NoErrorf(t, err, "%v", err)
			defer os.RemoveAll(dir)

			c := Config{}
			c.Segment.MaxStoreBytes = 32
			if scenario == "make new segment" {
				c.Segment.MaxIndexBytes = 13
			}
			log, err := NewLog(dir, c)
			assert.NoErrorf(t, err, "%v", err)

			fn(t, log)
		})

	}
}

func testAppendRead(t *testing.T, log *Log) {
	append := &api.Record{
		Value: []byte("hello world"),
	}

	off, err := log.Append(append)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), off)
	read, err := log.Read(off)
	assert.NoError(t, err)
	assert.Equal(t, append.Value, read.Value)
}

func testOutOfRangeError(t *testing.T, log *Log) {
	read, err := log.Read(1)
	assert.Nil(t, read)
	assert.Error(t, err)
}

func testInitExisting(t *testing.T, o *Log) {
	append := &api.Record{
		Value: []byte("hello world"),
	}
	for i := 0; i < 3; i++ {
		_, err := o.Append(append)
		assert.NoError(t, err)
	}
	assert.NoError(t, o.Close())

	off, err := o.LowestOffset()
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), off)
	off, err = o.HighestOffset()
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), off)

	n, err := NewLog(o.Dir, o.Config)
	assert.NoError(t, err)

	off, err = n.LowestOffset()
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), off)
	off, err = n.HighestOffset()
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), off)
}

func testReader(t *testing.T, log *Log) {
	append := &api.Record{
		Value: []byte("hello world"),
	}
	off, err := log.Append(append)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), off)

	reader := log.Reader()
	b, err := ioutil.ReadAll(reader)
	assert.NoError(t, err)

	read := &api.Record{}
	err = proto.Unmarshal(b[lenWidth:], read)
	assert.NoError(t, err)
	assert.Equal(t, append.Value, read.Value)
}

func testTruncate(t *testing.T, log *Log) {
	append := &api.Record{
		Value: []byte("hello world"),
	}
	for i := 0; i < 3; i++ {
		_, err := log.Append(append)
		assert.NoError(t, err)
	}

	_, err := log.Read(0)
	assert.NoError(t, err)

	err = log.Truncate(1)
	assert.NoError(t, err)

	_, err = log.Read(0)
	assert.NoError(t, err)
}
