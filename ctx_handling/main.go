package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func main() {
	// redis 에러난 경우 롤백을 해야한다.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		val, err := client.Get(ctx, "key").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer에서 키를 지운 후 확인", val)
	}()

	// 강제로 context cancel
	if err := client.Set(ctx, "key", "value", 0).Err(); err != nil {
		fmt.Println(err)
		cancel()
	}

	cancel()
}
