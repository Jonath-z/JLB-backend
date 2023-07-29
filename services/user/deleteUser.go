package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	if err := config.DB.Where("user_id = ?", userId).First(&entities.UserEntity{}).Delete(&entities.UserEntity{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully deleted",
	})
}
