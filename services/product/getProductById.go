package product

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductById(c *gin.Context) {
	productId := c.Param("id")

	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The product Id is required",
		})
		return
	}

	product := entities.ProductEntity{}
	result := config.DB.Where("product_id = ?", productId).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"error":   result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"product": product,
		})
	}
}
