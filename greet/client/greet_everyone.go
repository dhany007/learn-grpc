package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dhany007/learn-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %+v", err)
	}

	// create request yang akan dikirim
	reqs := []*pb.GreetRequest{
		{FirstName: "dhany"},
		{FirstName: "kalai"},
		{FirstName: "erna"},
	}

	waitc := make(chan struct{})

	// for send to server
	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %+v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		// close when done
		stream.CloseSend()
	}()

	// for receive response from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving response: %+v", err)
			}

			log.Printf("Recived: %+v", res.Result)
		}

		// close the channel when client finish
		close(waitc)
	}()

	// waiting for all goroutine close
	<-waitc
}
