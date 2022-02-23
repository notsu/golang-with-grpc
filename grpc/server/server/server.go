// Package server contains gRPC server handling
package server

import (
	"github.com/notsu/golang-with-grpc/grpc/server/proto"
)

// Server represents the gRPC server
type Server struct {
	proto.UnimplementedGreeterServer
}
