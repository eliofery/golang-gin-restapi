package main

import (
	"fmt"
	"github.com/eliofery/golang-restapi/database"
	"github.com/eliofery/golang-restapi/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	database.Init()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	fmt.Println("Старт сервера на порту 8080")
	err := server.Run(":8080")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось получить события",
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось получить ID события",
			"error":   err.Error(),
		})

		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось получить событие",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, event)
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
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось создать событие",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Событие успешно создано",
		"event":   event,
	})
}
