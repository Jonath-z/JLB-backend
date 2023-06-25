package main

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db"
	"github.com/Jonath-z/JLB-backend/services"
	"github.com/Jonath-z/JLB-backend/services/authentication"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvironmentVariables()
	config.ConnectDB()
	db.Migrate()
}

func main() {
	router := gin.Default()
	router.GET("/", services.DefaultRequest)
	router.POST("/signup", authentication.Sigup)

	router.Run()
}
