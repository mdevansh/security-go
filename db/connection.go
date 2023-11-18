package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri = os.Getenv("MONGO_URI")

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}
	return client, nil
}
