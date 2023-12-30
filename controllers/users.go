package controllers

import (
	"e_commerce_backend/helper"
	"e_commerce_backend/models"
	"e_commerce_backend/structs"
	"e_commerce_backend/utilities"
	"time"
)

func SignUp(data structs.UserSiginUp) (returnData utilities.ResponseJson) {
	User, err := models.GetUserByEmail(data.Email)
	if err != nil {
		utilities.ErrorResponse(&returnData, "User Already Present in Database")
		return
	}

	if User != nil && User.Email == data.Email {
		utilities.ErrorResponse(&returnData, "User email is present in the database")
		return
	}

	token, err := helper.CreateToken(data.FirstName, data.LastName, data.Email)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured while creating token")
		return
	}

	password, err := helper.HashPassword(data.Password)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Error occured at hashing password")
	}

	currentTimestamp := time.Now()

	if User == nil {
		id, err := models.AddUserDetails(&models.Users{
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Password:  password,
			UserToken: token,
			CreatedAt: currentTimestamp,
			UpdateAt:  currentTimestamp,
		})

		if err != nil {
			utilities.ErrorResponse(&returnData, "Error Ocuured while creating user")
			return
		}
		utilities.SuccessResponse(&returnData, id)
	}
	return
}

func SignIn(data structs.UserSignIn) (returnData utilities.ResponseJson) {
	userDetails, err := models.GetUserByEmail(data.Email)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Enter the correct email")
		return
	}

	check := helper.VerifyPassword(userDetails.Password, data.Password)

	if userDetails.Email == data.Email && check {
		utilities.SuccessResponse(&returnData, userDetails.UserToken)
		return
	} else {
		utilities.ErrorResponse(&returnData, "Email and Passwod does not match")
	}
	return
}
