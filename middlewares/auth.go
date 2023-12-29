package middlewares

import (
	"github.com/eliofery/golang-restapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Необходима авторизация",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if userId == 0 || err != nil {
		var errMsg any
		if err != nil {
			errMsg = err.Error()
		} else {
			errMsg = userId
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Необходима авторизация",
			"error":   errMsg,
		})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}
