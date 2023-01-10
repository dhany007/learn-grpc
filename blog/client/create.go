package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--- createBlog was invoked ---")

	blog := &pb.Blog{
		AuthorId: "Dhany",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("unexpected error: %+v", err)
	}

	log.Printf("Blog has beed created: %s\n", res.Id)

	return res.Id
}
