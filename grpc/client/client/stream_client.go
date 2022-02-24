package client

import (
	"context"
	"fmt"
	"log"
	"time"

	client "github.com/notsu/golang-with-grpc/grpc/client/proto"
)

// LotsOfGreetings sends streaming data to the gRPC server
func LotsOfGreetings(ctx context.Context, c client.GreeterClient) {
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	for i := 1; i < 10; i++ {
		stream.Send(&client.HelloRequest{
			Name: fmt.Sprintf("Notsu clone no.%d", i),
		})

		time.Sleep(3 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to close and receive: %s", err)
	}

	log.Printf("Response from server: %s", resp.Message)

	log.Println("Done!")
}
