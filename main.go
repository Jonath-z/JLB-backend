package main

import (
	"github.com/Jonath-z/JLB-backend/src/config"
	"github.com/Jonath-z/JLB-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDB()
	config.LoadEnv()
	router.GET("/", controllers.DefaultRequest)

	router.Run("localhost:4040")
}
