package database

import (
	"context"
	"fmt"
	"htf/src/utils"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(ctx context.Context, config *utils.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(config.RedisClusterURL)
	if err != nil {
		panic(err)
	}

	opt.OnConnect = func(ctx context.Context, c *redis.Conn) error {
		fmt.Println("Connected to redis...")
		return nil
	}

	client := redis.NewClient(opt)
	_, err = client.Ping(ctx).Result()
	if err != nil {
		err = fmt.Errorf("Utils.database.redis.NewRedisClient: %w", err)
		return nil, err
	}

	return client, nil

}
