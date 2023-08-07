package main

import (
	"github.com/Jonath-z/JLB-backend/config"
	"github.com/Jonath-z/JLB-backend/db"
	"github.com/Jonath-z/JLB-backend/services"
	"github.com/Jonath-z/JLB-backend/services/authentication"
	"github.com/Jonath-z/JLB-backend/services/product"
	"github.com/Jonath-z/JLB-backend/services/user"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvironmentVariables()
	config.ConnectDB()
	db.Migrate()
}

func main() {
	router := gin.Default()
	router.GET("/", services.DefaultRequest)
	router.PUT("/signup", authentication.Sigup)
	router.POST("/login", authentication.Login)
	router.GET("/users/:username", user.GetUserByUserName)
	router.GET("/users/get/:id", user.GetUserByUserId)
	router.PATCH("/users/update/:id", user.UpdateUser)
	router.GET("/users", user.GetAllUsers)
	router.DELETE("/users/:id", user.DeleteUser)
	router.DELETE("/users/products/:id", user.RemoveProduct)
	router.POST("/users/:id", user.SaveProduct)

	router.PUT("/product", product.AddProduct)
	router.GET("/products/:id", product.GetProductById)
	router.DELETE("/products/:id", product.DeleteProduct)
	router.GET("/products", product.GetAllProduct)
	router.PATCH("/products/:id", product.UpdateProduct)
	router.Run()
}
