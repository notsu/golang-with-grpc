package server

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/notsu/golang-with-grpc/grpc/server/proto"
)

// LotsOfReplies returns pong messages streaming back to the client
func (s *Server) LotsOfReplies(r *proto.HelloRequest, srv proto.Greeter_LotsOfRepliesServer) error {
	log.Printf("received: %v", r.Name)

	// use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(count int64) {
			defer wg.Done()

			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)

			resp := proto.HelloResponse{
				Message: fmt.Sprintf("Hello %s, I will send you a message no.#%d", r.Name, count),
			}

			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}

			log.Printf("Send a request no.%d back successfully", count)
		}(int64(i))
	}

	wg.Wait()

	return nil
}
