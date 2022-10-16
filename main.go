package main

import (
	"gin-simple-chat/config"
	"gin-simple-chat/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DropAllTables()
	config.MigrateDB()

	r := gin.Default()

	router.UserRouter(r)

	r.Run()
}
