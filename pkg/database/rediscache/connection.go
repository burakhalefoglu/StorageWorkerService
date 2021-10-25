package rediscache

import (
	"github.com/go-redis/redis/v8"
)

type Connection struct {
	Conn *redis.Client
}
var Conn = Connection{
	Conn: ConnectRedis(),
}

func ConnectRedis() *redis.Client {

	 return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}