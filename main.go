package main

import (
	"github.com/eliofery/golang-restapi/database"
	"github.com/eliofery/golang-restapi/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Init()

	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}
