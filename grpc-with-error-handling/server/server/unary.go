package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/notsu/golang-with-grpc/grpc-with-error-handling/server/proto"
)

// SayHello returns a pong message back to the client
func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Receive from the client name: %v", r.Name)

	st := status.New(codes.InvalidArgument, "Insufficient arguments")
	customErr := proto.Error{
		Code:    100,
		Message: "You may not send with sufficient arguments",
	}

	dt, err := st.WithDetails(&customErr)
	if err != nil {
		return nil, st.Err()
	}

	log.Printf("Send an error with code: %+v and also include a custom error\n", codes.InvalidArgument)

	return nil, dt.Err()
}
