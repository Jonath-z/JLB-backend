package authentication

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/Jonath-z/JLB-backend/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var loginData *models.LoginData

	err := c.BindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing password or email",
			"error":   err.Error(),
		})
		return
	}

	user := entities.UserEntity{}
	checkEmailErr := config.DB.Where("email = ?", loginData.Email).First(&user).Error
	if checkEmailErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Incorrect Email",
			"error":   checkEmailErr.Error(),
		})
	} else {
		isPasswordTrue := utilities.CheckPassword(loginData.Password, user.Password)
		if !isPasswordTrue {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Incorrect Password",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"user":    user,
			})
		}
	}
}
