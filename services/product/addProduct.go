package product

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db/entities"
	"github.com/Jonath-z/JLB-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AddProduct(c *gin.Context) {
	var newProduct *models.Product

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The product is not valid",
			"error":   err.Error(),
		})
		return
	}

	newProduct.ProductId = uuid.New().String()
	config.DB.Create(&entities.ProductEntity{
		Name:      newProduct.Name,
		Image:     newProduct.Image,
		Price:     newProduct.Price,
		ProductId: newProduct.ProductId,
	})

	c.JSON(http.StatusCreated, newProduct)
}
