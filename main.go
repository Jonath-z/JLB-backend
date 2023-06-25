package main

import (
	"log"

	"github.com/Jonath-z/JLB-backend/src/config"
	"github.com/Jonath-z/JLB-backend/src/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Cant not load environment variables")
	}
	
	router := gin.Default()
	config.ConnectDB()
	router.GET("/", controllers.DefaultRequest)

	router.Run()
}
