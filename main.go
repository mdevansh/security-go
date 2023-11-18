package main

import (
	"context"
	"fmt"
	"log"
	"security-go/api"
	"security-go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lol")

	client, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	router := gin.Default()

	api.InitRoutes(router, client)

	router.Run(":8080")

}
