package middleware

import (
	"fmt"
	"net/http"

	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckUserExist() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var newUser *models.User
		db := config.DB

		err := ctx.BindJSON(&newUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
		}

		fmt.Print(newUser)

		result := db.Where("name = ?", newUser.Username).First(&entities.UserEntity{})

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
			} else {
				ctx.JSON(http.StatusConflict, gin.H{
					"message": "username already exists",
				})
			}
		} else {
			ctx.Next()
		}
	})
}
