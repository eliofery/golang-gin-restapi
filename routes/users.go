package routes

import (
	"github.com/eliofery/golang-restapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось создать пользователя",
			"error":   err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось создать пользователя",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан",
		"user":    user,
	})
}
