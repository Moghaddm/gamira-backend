package configs

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{ClientName: "redis", Addr: addr, Password: password})

	return client
}
