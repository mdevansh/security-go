package api

import (
	"net/http"
	"os"
	"security-go/db"
	"security-go/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(router *gin.Engine, client *mongo.Client) {

	var dbName = os.Getenv("DB_NAME")
	var collectionName = os.Getenv("COLLECTION_NAME")

	collection := client.Database(dbName).Collection(collectionName)

	router.POST("/createUser", func(c *gin.Context) {

		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.CreateUser(collection, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failer to insert user " + newUser.Username + " :: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Successfully created user " + newUser.Username})
	})

}
