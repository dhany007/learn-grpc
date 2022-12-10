package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("function foLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "dhany"},
		{FirstName: "kalai"},
		{FirstName: "erna"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %+v", err)
	}

	// send to server
	for _, req := range reqs {
		log.Printf("sending request: %+v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second) // to see proces send to server
	}

	// close and get response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %+v", err)
	}

	log.Printf("LongGreet: %s", res.Result)
}
