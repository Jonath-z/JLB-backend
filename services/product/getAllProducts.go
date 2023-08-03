package product

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProduct(c *gin.Context) {
	var products []entities.ProductEntity

	result := config.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find the product",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Success",
		"products": products,
	})
}
