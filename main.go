package main

import (
	"fmt"
	"github.com/eliofery/golang-restapi/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	fmt.Println("Старт сервера на порту 8080")
	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}

func getEvent(ctx *gin.Context) {
	events := models.GetAllEvents()

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось создать событие",
			"error":   err.Error(),
		})

		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Событие успешно создано",
		"event":   event,
	})
}
