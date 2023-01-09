package main

import (
	"log"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	// use credentials
	tls := true // change to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("failed to load certificates: %+v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// end credentials

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %+v \n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c) // unary
	// doGreetManyTimes(c) // server stream
	// doLongGreet(c) // clien stream
	// doGreetEveryone(c) // bi-directional stream

	// doGreetWithDeadline(c, 4*time.Second) // success
	// doGreetWithDeadline(c, 2*time.Second) // deadline exceeded
}
