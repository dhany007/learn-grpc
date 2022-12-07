package main

import (
	"log"
	"net"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen tcp %+v \n", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve, %+v \n", err)
	}
}
