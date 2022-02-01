package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	db "todo-app-grpc/internal/services/DB"
	"todo-app-grpc/internal/services/blog"
	blogpb "todo-app-grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("starting server on port :50051...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("unable to listen on port :50051: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	mongoCtx := context.Background()
	db := db.New(mongoCtx)
	srv := blog.New(db.Database("mydb").Collection("blog"))

	blogpb.RegisterBlogServiceServer(s, srv)
	// Initialize MongoDb client
	fmt.Println("Connecting to mongoDB...")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve %v", err)
		}

	}()
	fmt.Println("Server succesfully starts on port: 50051")

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received

	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	fmt.Println("\nStopping the server... ")

	s.Stop()

	listener.Close()

	fmt.Println("Closing MongoDB conncetion")

	db.Disconnect(mongoCtx)

	fmt.Println("Done")
}
