package main

import (
	"log"

	pb "github.com/dhany007/learn-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %+v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c) // unary
	// doPrimes(c) // server streaming
	// doAverage(c) // client streaming
	// doMax(c)
	// doSqrt(c, 10)
	doSqrt(c, -10)
}
