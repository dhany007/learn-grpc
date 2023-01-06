package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("func doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Kalai",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("unexpected grpc error: %+v", err)
			}
		} else {
			log.Fatalf("a non grpc error: %+v", err)
		}
	}

	log.Printf("GreetWithDeadline: %+v\n", res.Result)
}
