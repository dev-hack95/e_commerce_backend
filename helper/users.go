package helper

import (
	"e_commerce_backend/structs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// func GenerateECDSAPrivateKey() (*ecdsa.PrivateKey, error) {
// 	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// }

var sampleSecretKey = []byte("SecretYouShouldHide")

func CreateToken(Firstname, Lastname, email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"first_name": Firstname,
			"last_name":  Lastname,
			"email":      email,
			"exp":        time.Now().Add(time.Hour * 24 * 45).Unix(),
			"admin":      true,
		})

	tokenString, err := token.SignedString(sampleSecretKey)
	fmt.Println(err)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "Error Occured while Hashing passsword", err
	}
	return string(hashPassword), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := true
	if err != nil {
		check = false
	}
	return check
}

func verifyToken(tokenString string) (*jwt.Token, *structs.JwtUserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.JwtUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, err
	}

	claims, ok := token.Claims.(*structs.JwtUserClaims)
	if !ok {
		return nil, nil, err
	}

	return token, claims, nil
}
