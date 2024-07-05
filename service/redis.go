package service

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()
var RDB *redis.Client

func InitRedis(addr string) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check if Redis is connected
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		zap.S().Fatal("Error connecting to Redis")
	}

	// Initialize the allowed keys, if they don't exist
	for key, value := range AllowedKeys {
		if RDB.Get(ctx, key).Val() == "" {
			err := RDB.Set(ctx, key, value, 0).Err()
			if err != nil {
				zap.S().Fatal("Error initializing key", zap.String("key", key), zap.Any("value", value))
			}
		}
	}
}
