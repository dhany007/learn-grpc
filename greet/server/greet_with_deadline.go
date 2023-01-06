package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with : %+v\n", in)

	// waiting 3 second for dealine
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
