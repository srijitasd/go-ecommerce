package config

import "github.com/redis/go-redis/v9"

func RedisInit() *redis.Client {
	var service *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return service
}
