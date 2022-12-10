package main

import (
	"context"
	"io"
	"log"
	"time"

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

// client streaming implementing
func doAverage(c pb.CalculatorServiceClient) {
	log.Printf("doAverage was invoked")

	numbers := []int32{1, 2, 3, 4, 5}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while receiving response from Average: %+v", err)
	}

	// send to server
	for _, number := range numbers {
		log.Printf("sending request: %+v\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})

		time.Sleep(1 * time.Second) // to see proccess send to server
	}

	// close and get response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %+v", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
