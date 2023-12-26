package main

import "fmt"

/*
main에서 먼저 종료
그 다음 작업 고루틴에서도 응답을 해준다. (오케이)
*/
func main() {
	c := make(chan string)
	quit := make(chan bool)
	// worker를 실행한다.
	go worker(c, quit)
	for i := 1; ; i++ {
		// 종료 시점을 설정
		if i%3 == 0 {
			quit <- true
			return
		}
		// work을 넣어준다.
		c <- fmt.Sprintf("work %d\n", i)
	}
}

func worker(c chan string, quit chan bool) { // Returns receive-only channel of strings.
	for {
		select {
		case msg := <-c:
			fmt.Printf("i got message: %s", msg)
		case <-quit:
			fmt.Printf("quit")
			return
		}
	}
}
