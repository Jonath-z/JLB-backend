package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveProduct(c *gin.Context) {
	userId := c.Param("id")
	var newProduct *models.Product

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The product is not valid",
			"error":   err.Error(),
		})
		return
	}

	user := entities.UserEntity{}
	queryResult := config.DB.Where("user_id = ?", userId).Preload("Products").First(&user)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"error":   queryResult.Error.Error(),
		})

		return
	}

	product := entities.ProductEntity{
		ProductId: newProduct.ProductId,
		Name:      newProduct.Name,
		Price:     newProduct.Price,
		Image:     newProduct.Image,
		UserId:    user.ID,
	}

	user.Products = append(user.Products, product)
	if err := config.DB.Where("user_id = ?", userId).Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save the products",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product saved successfully",
			"user":    user,
		})
	}
}
