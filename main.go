package main

import (
	"e_commerce_backend/handlers"
	"e_commerce_backend/utilities"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	utilities.EnableSQLDatabasesConfiguration()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	// Users
	router.POST("/v1/signup", handlers.SignUp)
	router.POST("/v1/signin", handlers.SignIn)
	// Sellers
	router.POST("/v1/seller/addproduct", handlers.AddProduct)
	router.GET("/v1/seller/getProductsById", handlers.GetProductById)
	// Search Functionalities
	router.GET("/v1/user/search", handlers.SearchProducts)
	router.GET("/v1/user/addProductToCart", handlers.AddProductToCart)

	router.Run()
}
