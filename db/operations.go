package db

import (
	"context"
	"security-go/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(collection *mongo.Collection, user models.User) error {
	_, err := collection.InsertOne(context.Background(), user)
	return err
}
