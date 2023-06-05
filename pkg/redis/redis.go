package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

type RSource struct {
	redisClient *redis.Client
}

func InitRedis(r *RSource) (*RSource, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	log.Printf("Connecting to Redis\n")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}
	return &RSource{
		redisClient: rdb,
	}, nil
}

type MessageRepository struct {
	redisClient *redis.Client
}

func NewMessageRepository(redisClient *redis.Client) *MessageRepository {
	return &MessageRepository{
		redisClient: redisClient,
	}
}

func RClose(r *RSource) error {
	if err := r.redisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis: %w", err)
	}
	return nil
}
