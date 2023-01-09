package main

import (
	"log"
	"net"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "localhost:50052"

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
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}
}
