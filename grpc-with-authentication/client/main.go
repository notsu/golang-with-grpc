package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/notsu/golang-with-grpc/grpc-with-authentication/client/data"
	client "github.com/notsu/golang-with-grpc/grpc-with-authentication/client/proto"
)

func main() {
	fmt.Println("Run a gRPC client")

	// Timeout in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	perRPC := oauth.NewOauthAccess(fetchToken())

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.NewOauthAccess requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewGreeterClient(conn)

	sayHello(ctx, c)
}

func sayHello(ctx context.Context, c client.GreeterClient) {
	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "Notsu"})
	if err != nil {
		log.Printf("Error from the server: %s", err)

		return
	}

	log.Printf("Response from server: %s", response.Message)
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
