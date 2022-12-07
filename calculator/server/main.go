package main

import (
	"log"
	"net"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %+v\n", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}
}
