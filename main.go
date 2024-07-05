package main

import (
	"net"
	"os"

	"github.com/Fan-Fuse/config-service/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		zap.S().Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterServer(s)

	zap.S().Info("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		zap.S().Fatalf("Failed to serve: %v", err)
	}
}
