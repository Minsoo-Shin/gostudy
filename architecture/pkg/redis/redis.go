package redis

import "github.com/redis/go-redis/v9"

const (
	DailyTestKey = "daily-test:%v:%v" //daily-test:{teacherID}:{testID}
	Nil          = redis.Nil
)

type Client = redis.Client

func New() *Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "myredis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return c
}
