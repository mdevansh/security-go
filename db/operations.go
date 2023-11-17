package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(collection *mongo.Collection, user interface{}) error {
	_, err := collection.InsertOne(context.Background(), user)
	return err
}
