package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "dhany",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while stream: %+v", err)
	}

	for {
		msg, err := stream.Recv() // stream received

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading the stream: %+v", err)
		}

		log.Printf("GreetManyTimes: %+v\n", msg.Result)
	}
}
