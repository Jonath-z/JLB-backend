package authentication

import (
	"net/http"

	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/Jonath-z/JLB-backend/utilities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Sigup(c *gin.Context) {
	var newUser *models.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
	}

	newUser.UserId = uuid.New().String()
	hashedPassword, _ := utilities.HashPassword(newUser.Password)
	newUser.Password = hashedPassword

	config.DB.Create(&entities.UserEntity{
		Name:     &newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	c.JSON(http.StatusCreated, newUser)
}
