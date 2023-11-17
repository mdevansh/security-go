package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://username:password@localhost:27017/?authSource=mydatabase")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}
	return client, nil
}
