package handlers

import (
	"e_commerce_backend/controllers"
	"e_commerce_backend/structs"
	"e_commerce_backend/utilities"
	"encoding/json"
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

	switch {
	case !utilities.IsEmpty(inputobj):
		returnData = controllers.SignIn(inputobj)
	default:
		utilities.ErrorResponse(&returnData, "Unbale to login")
	}

	c.JSON(returnData.Code, returnData)
}
