package product

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	product := entities.ProductEntity{}
	result := config.DB.Where("product_id = ?", productId).First(&product).Delete(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully deleted",
	})
}
