package handlers

import (
	"e_commerce_backend/controllers"
	"e_commerce_backend/structs"
	"e_commerce_backend/utilities"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect posts data")
		c.JSON(returnData.Code, returnData)
		return
	}

	inputobj := structs.UserSiginUp{}
	err = json.Unmarshal(body, &inputobj)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failed to unmarshal JSON")
		c.JSON(returnData.Code, returnData)
		return
	}

	switch {
	case !utilities.IsEmpty(inputobj):
		returnData = controllers.SignUp(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Something Went Wrong")
	}

	c.JSON(returnData.Code, returnData)
}

func SignIn(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: Unable to read body")
		c.JSON(returnData.Code, returnData)
		return
	}

	inputobj := structs.UserSignIn{}
	_ = json.Unmarshal(body, &inputobj)

	fmt.Println(inputobj)

	switch {
	case !utilities.IsEmpty(inputobj):
		returnData = controllers.SignIn(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Unbale to login")
	}

	c.JSON(returnData.Code, returnData)
}

func AddProduct(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: Unable to read body")
		c.JSON(returnData.Code, returnData)
		return
	}

	inputobj := structs.Product{}
	err = json.Unmarshal(body, &inputobj)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failed to unmarshal JSON")
		c.JSON(returnData.Code, returnData)
		return
	}

	switch {
	case inputobj.ProductName != "":
		returnData = controllers.AddProduct(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Error occurred while adding product")
	}
	c.JSON(returnData.Code, returnData)
}

func GetProductById(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	productId := c.Request.URL.Query().Get("product_id")

	switch {
	case !utilities.IsEmpty(productId):
		returnData = controllers.GetProductById(productId)
	default:
		utilities.ErrorResponse(&returnData, "Error occured while fetching data from database")
	}
	return
}
