package dao

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	r *redis.Client
}

func NewRedisClient() *RedisClient {
	//连接服务器
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	r := RedisClient{r: db}
	return &r
}

func (r *RedisClient) Set(key string, value interface{}) (err error) {
	result := r.r.Set(key, value, time.Minute*10)
	return result.Err()
}

func (r *RedisClient) Get(key string) (value string, err error) {
	result := r.r.Get(key)
	return result.Result()
}

func (r *RedisClient) MemoryUsage(key string, samples ...int) {
	r.r.MemoryUsage(key, samples...)
}
