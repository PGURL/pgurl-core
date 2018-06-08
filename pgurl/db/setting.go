package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/PGURL/pgurl-core/pgurl/config"
)

var redis_client *redis.Client = RedisCon()

func  RedisCon () *redis.Client{
	config := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),  // use default DB
	})
    return client
}
