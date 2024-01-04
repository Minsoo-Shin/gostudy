package main

import "fmt"

/*
main에서 먼저 종료
그 다음 작업 고루틴에서도 응답을 해준다. (오케이)
*/
func main() {
	/*
		## worker
		worker는 goroutine 비동기로 channel을 subscribe하고 있는다.
		main 함수에서 일을 주고 종료 시점에서
		quit <- true를 넣어서 go routine을 종료해준다.
	*/
	//c := make(chan string)
	//quit := make(chan bool)
	//// worker를 실행한다.
	//go worker(c, quit)
	//for i := 1; ; i++ {
	//	// 종료 시점을 설정
	//	if i%10 == 0 {
	//		quit <- false
	//		return
	//	}
	//	// work을 넣어준다.
	//	c <- fmt.Sprintf("work %d\n", i)
	//}

	/*
		goroutine간 통신을 위해서 chan을 사용했다. 단방향뿐만 아니고 양방향 통신까지 가능하다.
	*/
	c := make(chan string)
	quit := make(chan string)
	// worker를 실행한다.
	go worker1(c, quit)
	for i := 1; ; i++ {
		// 종료 시점을 설정
		if i%10 == 0 {
			quit <- "main says: good bye\n"
			fmt.Println(<-quit)
			return
		}
		// work을 넣어준다.
		c <- fmt.Sprintf("work %d\n", i)
	}
}

func worker1(c chan string, quit chan string) { // Returns receive-only channel of strings.
	for {
		select {
		case msg := <-c:
			fmt.Printf("i got message: %s", msg)
		case goodbyemsg := <-quit:
			fmt.Printf("%v", goodbyemsg)
			quit <- "worker1 says: see you\n"
		}
	}
}

func worker(c chan string, quit chan bool) { // Returns receive-only channel of strings.
	for {
		select {
		case msg := <-c:
			fmt.Printf("i got message: %s", msg)
		case boolean := <-quit:
			if boolean == true {
				fmt.Println("quit for true")
				return
			} else {
				fmt.Println("not quit for false")
				return
			}
		}
	}
}
