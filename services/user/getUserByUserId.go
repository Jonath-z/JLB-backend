package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByUserId(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}

	user := entities.UserEntity{}
	err := config.DB.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user id not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}
