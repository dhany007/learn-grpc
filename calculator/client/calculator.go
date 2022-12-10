package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
)

// unary implementing
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

// server streaming implementing
func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes function was invoked")
	req := &pb.PrimeRequest{
		Number: 123513456,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while stram: %+v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading the stream: %+v", err)
		}

		log.Printf("Primes = %d\n", msg.Result)
	}
}
