package product

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The product id is required",
		})

		return
	}

	updatedProduct := &Product{}
	if err := c.ShouldBindJSON(updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	updateResult := config.DB.Where("product_id = ?", productId).First(&entities.ProductEntity{}).Updates(entities.ProductEntity{
		Name:  updatedProduct.Name,
		Image: updatedProduct.Image,
		Price: updatedProduct.Price,
	})
	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update the product",
			"error":   updateResult.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated successfully",
		})
	}
}
