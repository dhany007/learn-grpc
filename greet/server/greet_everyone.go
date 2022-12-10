package main

import (
	"io"
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	// catch all of request from client
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while create stream: %+v", err)
		}

		res := "Hello " + req.FirstName + "!"

		// send stream to the client
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while send data to client: %+v", err)
		}
	}
}
