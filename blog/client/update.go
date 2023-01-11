package main

import (
	"context"
	"log"

	pb "github.com/dhany007/learn-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Kalai",
		Title:    "New Title blog",
		Content:  "Content of the blog, with awesome additions",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("error while updating: %+v", err)
	}

	log.Println("Blog has been updated")
}
