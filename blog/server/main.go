package main

import (
	"context"
	"log"
	"net"

	pb "github.com/dhany007/learn-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	addr       string = "localhost:50052"
	collection *mongo.Collection
)

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// start connection to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")
	// end

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed ti serve: %+v\n", err)
	}
}
