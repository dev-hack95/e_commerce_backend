package structs

import "github.com/golang-jwt/jwt/v5"

type UserSiginUp struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtUserClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
	jwt.Claims
}
