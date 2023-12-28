package routes

import (
	"github.com/eliofery/golang-restapi/models"
	"github.com/eliofery/golang-restapi/utils"
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

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Не удалось войти",
			"error":   err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Не удалось войти",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Не удалось войти",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Вы вошли в систему",
		"token":   token,
	})
}
