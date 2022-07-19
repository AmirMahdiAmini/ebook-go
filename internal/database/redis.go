package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func SetupRedis() *redis.Client {
	fmt.Println("REDIS : ", os.Getenv("REDISDB_URI"))
	option, err := redis.ParseURL(os.Getenv("REDISDB_URI"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(option)
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     os.Getenv("REDIS_ADDRESS"),
	// 	Password: os.Getenv("REDIS_PASSWORD"),
	// 	DB:       0,
	// })
	_, err = client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return client
}
