package service

import (
	"context"

	"github.com/Fan-Fuse/config-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	proto.UnimplementedConfigServiceServer
	subscribers map[string][]chan *proto.SubscribeResponse
}

// NewServer creates a new ConfigServiceServer.
func RegisterServer(s *grpc.Server) {
	proto.RegisterConfigServiceServer(s, &server{
		subscribers: make(map[string][]chan *proto.SubscribeResponse),
	})
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

func (s *server) SetKey(ctx context.Context, in *proto.SetKeyRequest) (*emptypb.Empty, error) {
	// TODO: Check if the key is allowed
	// Set the value for the key
	err := RDB.Set(ctx, in.Key, in.Value, 0).Err()
	if err != nil {
		return nil, err
	}

	// Notify subscribers
	for _, ch := range s.subscribers[in.Key] {
		ch <- &proto.SubscribeResponse{
			Key:   in.Key,
			Value: in.Value,
		}
	}

	// Return the key and value as confirmation
	return &emptypb.Empty{}, nil
}

func (s *server) Subscribe(in *proto.SubscribeRequest, stream proto.ConfigService_SubscribeServer) error {
	ch := make(chan *proto.SubscribeResponse)
	for _, key := range in.Keys {
		s.subscribers[key] = append(s.subscribers[key], ch)
	}

	// Listen for subscription updates
	for update := range ch {
		if err := stream.Send(update); err != nil {
			return err
		}
	}
	return nil
}
