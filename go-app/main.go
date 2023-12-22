package main

import (
	"net/http"

	"github.com/caoquy2000/meeting-app/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	router.Run(":8000")
}
