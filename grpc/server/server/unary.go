package server

import (
	"context"
	"fmt"
	"log"

	"github.com/notsu/golang-with-grpc/grpc/server/proto"
)

// SayHello returns a pong message back to the client
func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Receive from the client name: %v", r.Name)

	return &proto.HelloResponse{
		Message: fmt.Sprintf("Hello %s, we receive your message", r.Name),
	}, nil
}
