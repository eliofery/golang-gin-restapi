package routes

import (
	"github.com/eliofery/golang-gin-restapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
		return
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

	event.UserID = ctx.GetInt("userId")
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось создать событие",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Событие успешно создано",
		"event":   event,
	})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось получить ID события",
			"error":   err.Error(),
		})
		return
	}

	userId := ctx.GetInt("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось получить событие",
			"error":   err.Error(),
		})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Недостаточно прав для обновления события",
		})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось обновить событие",
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось обновить событие",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Событие успешно обновлено",
		"event":   updatedEvent,
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось получить ID события",
			"error":   err.Error(),
		})
		return
	}

	userId := ctx.GetInt("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось получить событие",
			"error":   err.Error(),
		})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Недостаточно прав для удаления события",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось удалить событие",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Событие успешно удалено",
	})
}
