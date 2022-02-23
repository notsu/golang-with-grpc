package server

import (
	"io"
	"log"
	"time"

	"github.com/notsu/golang-with-grpc/grpc/server/proto"
)

// LotsOfGreetings receives streaming greetings and send back a pong message
func (s *Server) LotsOfGreetings(srv proto.Greeter_LotsOfGreetingsServer) error {
	i := 1

	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Receive from client: %s", in.Name)

		time.Sleep(1 * time.Second)

		if i > 5 {
			out := &proto.HelloResponse{
				Message: "OK, I'm done",
			}

			if err := srv.SendAndClose(out); err != nil {
				return err
			}

			return nil
		}

		i++
	}
}
