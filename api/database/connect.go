package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx = context.TODO()

func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = client.Database("investbot")
	fmt.Println("Connected to MongoDB!")
	return client, nil
}
