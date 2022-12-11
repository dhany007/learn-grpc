package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// uanry implementing
func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with: %+v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

// server streaming implementing
func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes was invoked with: %+v\n", in)

	var (
		divisor int32 = 2
		number  int32 = in.Number
	)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number = number / divisor
		} else {
			divisor += 1
		}
	}

	return nil
}

// client streaming api implementing
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")

	var (
		sum     int32 = 0
		counter int32 = 0
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("error while receive stream: %+v", err)
		}

		log.Printf("Receiving: %d\n", req.Number)

		sum += req.Number
		counter += 1
	}
}

// bi-directional streaming api implementing
func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var (
		max int32 = int32(math.MinInt32) // to get minimum number of int32
	)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while receiving stream: %+v", err)
		}

		if req.Number > max {
			max = req.Number
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("error while send data to client: %+v", err)
			}
		}
	}
}

// Error Handling
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %+v", in)

	number := in.Number

	// handle when number is negative
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received a negative number: %d", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
