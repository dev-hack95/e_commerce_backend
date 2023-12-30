package handlers

import (
	"e_commerce_backend/utilities"

	"github.com/gin-gonic/gin"
)

func AddProductToCart(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	productId := c.Request.URL.Query().Get("product_id")
	_, _ = returnData, productId
	return
}
