package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("function LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("while receive stream: %+v", err)
		}

		log.Printf("Reveiveing: %+v\n", req)
		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
