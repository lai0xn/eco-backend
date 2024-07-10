package redis

import (
	"context"
	"os"

	"github.com/lai0xn/squid-tech/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.TODO()
var client *redis.Client

func Connect() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Username: os.Getenv("REDIS_USERNAME"),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Err(err)
	}
	logger.Logger.Info().Msg("Connected to Redis successfully!")
}

func GetClient() *redis.Client {
	return client
}
