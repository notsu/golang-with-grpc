package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	client "github.com/notsu/golang-with-grpc/grpc/client/proto"
)

// BidiHello streams both between the client and server
func BidiHello(ctx context.Context, c client.GreeterClient) {
	done := make(chan bool)

	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	go func(stream client.Greeter_BidiHelloClient) {
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

	go func(stream client.Greeter_BidiHelloClient) {
		i := 1

		for {
			name := fmt.Sprintf("Notsu clone no.%d", i)

			err := stream.Send(&client.HelloRequest{
				Name: name,
			})
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error when receive a message: %s", err)
			}

			log.Printf("Send a message to server with name: %s", name)

			i++

			time.Sleep(1 * time.Second)

			if i > 10 {
				if err := stream.CloseSend(); err != nil {
					log.Fatalf("Error when trying to close send: %s", err)
				}

				log.Println("OK, I'm done")

				break
			}
		}
	}(stream)

	<-done

	log.Println("Done!")
}
