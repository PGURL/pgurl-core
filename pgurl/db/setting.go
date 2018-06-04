package db

import "github.com/go-redis/redis"

var redis_client *redis.Client = RedisCon()

func  RedisCon () *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.31:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
    return client
}
