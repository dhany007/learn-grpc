package main

import (
	"context"
	"io"
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

// client streaming api implementing
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")

	var (
		avg     float32 = 0
		sum     float32 = 0
		counter int32   = 0
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalf("error while receive stream: %+v", err)
		}

		log.Printf("Receiving: %d\n", req.Number)

		sum += float32(req.Number)
		counter += 1

		avg = sum / float32(counter)
	}

}
