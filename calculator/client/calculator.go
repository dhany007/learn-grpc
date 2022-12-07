package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum function was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("could not sum: %+v\n", err)
	}

	log.Printf("SUM: %d", res.Result)
}
