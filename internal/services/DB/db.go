package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context) *mongo.Client {
	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to connect to mongoDB: %v\n", err)
	} else {
		fmt.Println("Connceted to mondoDB")
	}

	return db

}
