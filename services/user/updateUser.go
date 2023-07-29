package user

import (
	"github.com/Jonath-z/JLB-backend/db/entities"
	"net/http"

	"github.com/Jonath-z/JLB-backend/config"
	"github.com/gin-gonic/gin"
)

type updatUserRequest struct {
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	ProfileThumbnail string `json:"profileThumbnail,omitempty"`
	ProfileUrl       string `json:"profileUrl,omitempty"`
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	var updatedUser = &updatUserRequest{}

	if err := c.BindJSON(updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"err":     err.Error(),
		})
		return
	}

	updateResult := config.DB.Where("user_id = ?", userId).First(&entities.UserEntity{}).Updates(entities.UserEntity{
		Name:             &updatedUser.Name,
		Email:            updatedUser.Email,
		ProfileThumbnail: &updatedUser.ProfileThumbnail,
		ProfileUrl:       &updatedUser.ProfileUrl,
	})

	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update user",
			"error":   updateResult.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated successfully",
		})
	}
}
