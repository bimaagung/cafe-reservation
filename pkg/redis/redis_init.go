package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
)

func NewRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,  // use default DB
	})
	rdb.AddHook(apmgoredis.NewHook())
	return rdb
}