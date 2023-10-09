package main

import (
	"fmt"
	"time"
)

func main() {
	//// 정수형 채널을 생성한다
	//ch1 := make(chan bool)
	//
	//go run1(ch1)
	//
	//retry := make(chan string)
	//for {
	//	select {
	//	case res := <-ch1:
	//		fmt.Println("run1", res)
	//		makeRetry(retry)
	//	case re := <-retry:
	//		fmt.Println(re)
	//	}
	//}
	//
	///*
	//	반복적으로 시간을 프린트한다.
	//	외부에서 종료 시그널을 주면 go routine을 종료한다.
	//*/
	//
	//infiniteLoop(terminate)

}

func infiniteLoop(terminate <-chan time.Time) {
	for {
		select {
		case <-terminate:
			fmt.Println("go routine exit")
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}

	}
}
