package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByUserName(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing username",
		})
		return
	}

	user := entities.UserEntity{}
	err := config.DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "username not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}
