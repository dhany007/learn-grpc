package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("error happening while reading: %+v", err)
		return nil
	}

	log.Printf("blog was read: %+v\n", res)

	return res
}
