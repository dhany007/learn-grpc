package main

import (
	"fmt"
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) (err error) {
	log.Printf("GreetManyTimes was invoked with: %+v", in)

	// we will make 10 times
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
