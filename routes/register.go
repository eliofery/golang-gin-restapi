package routes

import (
	"github.com/eliofery/golang-gin-restapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
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

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось зарегистрироваться на событие",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Вы успешно зарегистрировались на событие",
		"event":   event,
	})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось получить ID события",
			"error":   err.Error(),
		})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось отменить регистрацию на событие",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Вы отменили регистрацию на событие",
		"event":   event,
	})
}
