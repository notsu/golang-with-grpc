package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcClient "github.com/notsu/golang-with-grpc/grpc/client/client"
	client "github.com/notsu/golang-with-grpc/grpc/client/proto"
)

var (
	method = "bidiHello"
)

func main() {
	fmt.Println("Run a gRPC client")
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewGreeterClient(conn)

	switch method {
	case "sayHello":
		grpcClient.SayHello(ctx, c)
	case "lotsOfReplies":
		grpcClient.LotsOfReplies(ctx, c)
	case "lotsOfGreetings":
		grpcClient.LotsOfGreetings(ctx, c)
	case "bidiHello":
		grpcClient.BidiHello(ctx, c)
	}
}
