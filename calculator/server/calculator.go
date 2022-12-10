package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
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
