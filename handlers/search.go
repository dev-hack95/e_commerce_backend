package handlers

import (
	"e_commerce_backend/controllers"
	"e_commerce_backend/utilities"

	"github.com/gin-gonic/gin"
)

func SearchProducts(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	text := c.Request.URL.Query().Get("text")

	switch {
	case !utilities.IsEmpty(text):
		returnData = controllers.SearchProducts(text)
	default:
		utilities.ErrorResponse(&returnData, "Error occured while searching product pls try again!")
	}
	return
}
