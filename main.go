package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	fmt.Println("Старт сервера на порту 8080")
	server.Run(":8080")
}
