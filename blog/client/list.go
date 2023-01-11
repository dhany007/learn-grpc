package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dhany007/learn-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---listBlogs was invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling ListBlogs: %+v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("something error: %+v", err)
		}

		log.Println(res)
	}
}
