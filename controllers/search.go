package controllers

import (
	"e_commerce_backend/models"
	"e_commerce_backend/utilities"
)

func SearchProducts(text string) (returnData utilities.ResponseJson) {
	resultData, err := models.SearchProducts(text)

	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured while searchin text in database!")
		return
	}

	utilities.SuccessResponse(&returnData, resultData)
	return
}
