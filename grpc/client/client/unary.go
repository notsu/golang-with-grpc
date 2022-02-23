package client

import (
	"context"
	"log"

	client "github.com/notsu/golang-with-grpc/grpc/client/proto"
)

// SayHello sends a unary call to the gRPC server
func SayHello(ctx context.Context, c client.GreeterClient) {
	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "Notsu"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
