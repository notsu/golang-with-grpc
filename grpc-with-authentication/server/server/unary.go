package server

import (
	"context"
	"fmt"
	"log"

	"github.com/notsu/golang-with-grpc/grpc-with-authentication/server/proto"
)

// SayHello returns a pong message back to the client
func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Recieve from the client name: %v", r.Name)

	return &proto.HelloResponse{
		Message: fmt.Sprintf("Hello %s, how are you?", r.Name),
	}, nil
}
