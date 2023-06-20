package server

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestLog_Append(t *testing.T) {
	// 여러개의 Record Append (go routine)
	log := NewLog()

	n := 100000
	wait := sync.WaitGroup{}
	wait.Add(n)
	for i := 0; i < n; i++ {
		num := i
		go func() {
			defer wait.Done()
			_, err := log.Append(Record{
				Value: []byte(strconv.Itoa(num)),
			})
			if err != nil {
				t.Errorf("%v", err)
			}
		}()
	}

	wait.Wait()

	// Record누락된건 없는지 확인
	if len(log.records) != n {
		fmt.Println(len(log.records))
		t.Errorf("length of records: %v", len(log.records))
	}
}
