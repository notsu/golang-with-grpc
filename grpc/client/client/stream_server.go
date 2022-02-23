package client

import (
	"context"
	"io"
	"log"

	client "github.com/notsu/golang-with-grpc/grpc/client/proto"
)

// LotsOfReplies sends a gRPC call and receive streaming data back
func LotsOfReplies(ctx context.Context, c client.GreeterClient) {
	done := make(chan bool)

	stream, err := c.LotsOfReplies(ctx, &client.HelloRequest{Name: "Notsu"})
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	go func(stream client.Greeter_LotsOfRepliesClient) {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error when receive a message: %s", err)
			}

			log.Printf("Response from server: %s", msg.Message)
		}
	}(stream)

	<-done

	log.Println("Done!")
}
