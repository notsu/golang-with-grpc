package server

import (
	"github.com/notsu/golang-with-grpc/grpc-with-error-handling/server/proto"
)

// Server represents the gRPC server
type Server struct {
	proto.UnimplementedGreeterServer
}
