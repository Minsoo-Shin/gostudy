package redis

import "github.com/redis/go-redis/v9"

const (
	Nil = redis.Nil
)

func New() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "myredis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return c
}
