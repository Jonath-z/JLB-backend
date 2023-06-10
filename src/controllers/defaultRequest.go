package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func DefaultRequest(c *gin.Context) {

	c.JSON(200, gin.H{
		"message":   "welcome to JLB backend system",
		"postgress": os.Getenv("POSTGRES_HOST"),
	})
}
