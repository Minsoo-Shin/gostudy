package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// loop
	for i := 0; i < 100; i++ {
		go func() {
			GetHttp()
		}()
	}
}

func GetHttp() {
	// GET 호출
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
