package client

import (
	"context"
	"log"

	"google.golang.org/grpc/status"

	client "github.com/notsu/golang-with-grpc/grpc-with-error-handling/client/proto"
)

// SayHello sends a unary call to the gRPC server
func SayHello(ctx context.Context, c client.GreeterClient) {
	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "Notsu"})
	if err != nil {
		s := status.Convert(err)

		log.Printf("Status: %s", s.Code())

		for _, d := range s.Details() {
			switch info := d.(type) {
			case *client.Error:
				log.Printf("Custom failure: %s", info)
			default:
				log.Fatalf("Error when calling SayHello: %s", err)
			}
		}

		return
	}

	log.Printf("Response from server: %s", response.Message)
}
