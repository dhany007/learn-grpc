package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet function was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "dhany",
	})

	if err != nil {
		log.Fatalf("could not greet: %+v\n", err)
	}

	log.Printf("Greetings: %s\n", res.Result)
}
