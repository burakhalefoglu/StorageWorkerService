package rediscache

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
)

type Connection struct {
	Conn *redis.Client
}
var Conn = Connection{
	 Conn: ConnectRedis(),
}

func ConnectRedis() *redis.Client{
	godotenv.Load()
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONN"),
		Password:  os.Getenv("REDIS_PASS"),
		DB:       0,
	})

}