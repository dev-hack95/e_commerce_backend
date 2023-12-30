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

	router.Run()
}
