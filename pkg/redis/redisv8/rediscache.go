package rediscachev8

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)


type redisCache struct {
	Client *redis.Client
}

func RedisCacheConstructor() *redisCache {
	return &redisCache{Client: getClient()}
}

func getClient() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONN"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (r *redisCache) Get(key string) (map[string]string, error) {

	result := r.Client.HGetAll(context.Background(),key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (r *redisCache) Add(key string, value map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(),key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) Delete(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(),key, fields ...)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteAll(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(),key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}
