package db

import "fmt"

func Ping() string {
	pong, err := redis_client.Ping().Result()
	fmt.Println(pong, err)
	return pong
}
