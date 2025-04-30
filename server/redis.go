package server

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func ConnectRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
