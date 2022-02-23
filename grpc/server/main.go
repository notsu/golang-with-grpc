package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/notsu/golang-with-grpc/grpc/server/proto"
	"github.com/notsu/golang-with-grpc/grpc/server/server"
)

func main() {
	fmt.Println("Run a gRPC server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGreeterServer(grpcServer, &server.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
