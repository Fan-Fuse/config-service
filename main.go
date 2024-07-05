package main

import (
	"os"

	"github.com/Fan-Fuse/config-service/service"
	"go.uber.org/zap"
)

func init() {
	// Read some values from environment (Redis address, etc.)
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}

	// Initialize logger
	logger := zap.Must(zap.NewProduction())

	zap.ReplaceGlobals(logger)

	zap.S().Info("Service initializing...")

	// Connect to Redis
	service.InitRedis(redisAddress)
}

func main() {
	zap.S().Info("Service starting...")
}
