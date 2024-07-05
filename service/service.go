package service

import (
	"context"

	"github.com/Fan-Fuse/config-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedConfigServiceServer
}

// NewServer creates a new ConfigServiceServer.
func RegisterServer(s *grpc.Server) {
	proto.RegisterConfigServiceServer(s, &server{})
}

func (s *server) GetKey(ctx context.Context, in *proto.GetKeyRequest) (*proto.GetKeyResponse, error) {
	// TODO: Check if the key is allowed
	// Get the value for the key
	val := RDB.Get(ctx, in.Key).Val()
	return &proto.GetKeyResponse{
		Key:   in.Key,
		Value: val,
	}, nil
}
