package my_redis

import "github.com/go-redis/redis"

type MyRedis struct {
}

func NewMyRedis() *MyRedis {
	return &MyRedis{}
}

func (receiver *MyRedis) Connect() *redis.Client {
	addr := "43.138.174.42:6379"
	password := "123456"
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		PoolSize: 2,
	})
}
