package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	var updatedUser *models.User

	err := c.BindJSON(updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"err":     err,
		})

		return
	}

	dbErr := config.DB.Where("user_id = ?", updatedUser.UserId).Updates(models.User{
		Username:         updatedUser.Username,
		Email:            updatedUser.Email,
		ProfileThumbnail: updatedUser.ProfileThumbnail,
		ProfileUrl:       updatedUser.ProfileUrl,
	}).Error
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated successfully",
		})
	}
}
