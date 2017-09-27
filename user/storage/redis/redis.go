package redis

import (
	"os"

	goredis "github.com/go-redis/redis"
)

var Client *RedisStorage

type RedisStorage struct {
	*goredis.Client
}

func init() {
	// RedisClient
	addr := os.Getenv("REDIS_ADDR")
	passwd := os.Getenv("REDIS_PASSWD")
	if addr == "" {
		addr = "localhost:6379"
	}
	rc := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	// RedisStorage
	Client = &RedisStorage{
		Client: rc,
	}
}
