package main

import (
	"log"
	"time"

	pb "github.com/dhany007/learn-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %+v \n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c) // unary
	// doGreetManyTimes(c) // server stream
	// doLongGreet(c) // clien stream
	// doGreetEveryone(c) // bi-directional stream

	// doGreetWithDeadline(c, 4*time.Second) // success
	doGreetWithDeadline(c, 2*time.Second) // deadline exceeded

}
