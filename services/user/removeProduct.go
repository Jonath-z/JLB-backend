package user

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveProductRequestBody struct {
	ProductId string
}

func RemoveProduct(c *gin.Context) {
	userId := c.Param("id")
	var productIdToRemove RemoveProductRequestBody

	err := c.ShouldBindJSON(&productIdToRemove)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product id is required",
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

	var updatedProduct []entities.ProductEntity
	for _, product := range user.Products {
		if product.ProductId != productIdToRemove.ProductId {
			updatedProduct = append(updatedProduct, product)
		}
	}

	user.Products = updatedProduct
	if err := config.DB.Where("user_id = ?", userId).Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete the product",
			"error":   err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
			"user":    user,
		})
	}
}
