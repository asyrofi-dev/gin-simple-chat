package main

import (
	"gin-simple-chat/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Gin Simple Chat",
		})
	})
	r.Run()
}
