package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "43.138.174.42:6379",
		Password: "123456",
		PoolSize: 2,
	})

	result, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}

	result, err = client.Get("app_id").Result()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

	result, err = client.Get("secret").Result()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

}
