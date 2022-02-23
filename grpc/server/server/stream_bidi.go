package server

import (
	"fmt"
	"io"
	"log"

	"github.com/notsu/golang-with-grpc/grpc/server/proto"
)

// BidiHello steams greeting for both client and server
func (s *Server) BidiHello(srv proto.Greeter_BidiHelloServer) error {
	i := 1

	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Receive a name from the client: %s", in.Name)

		out := &proto.HelloResponse{
			Message: fmt.Sprintf("Hello %s, how are you today?", in.Name),
		}

		if err := srv.Send(out); err != nil {
			return err
		}

		log.Printf("Send a message back to the client name: %s", in.Name)

		i++
	}
}
