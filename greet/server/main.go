package main

import (
	"log"
	"net"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen tcp %+v \n", err)
	}

	log.Printf("listening on %s", addr)

	// use credentials
	tls := true // change to false if needed
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to load certificates: %+v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}
	// end credentials

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve, %+v \n", err)
	}
}
