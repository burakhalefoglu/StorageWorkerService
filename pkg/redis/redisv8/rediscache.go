package rediscachev8

import (
	"StorageWorkerService/pkg/helper"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type redisCache struct {
	Client *redis.Client
}

func RedisCacheConstructor() *redisCache {
	return &redisCache{Client: getClient()}
}

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     helper.ResolvePath("REDIS_HOST", "REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,                       // use default DB
	})
	func() {
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("RedisConnection", "ConnectRedis", err)
		}
	}()

	return client
}

func (r *redisCache) Get(key string) (map[string]string, error) {

	result := r.Client.HGetAll(context.Background(), key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (r *redisCache) Add(key string, value map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(), key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) Delete(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(), key, fields...)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteAll(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}
