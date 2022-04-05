package model

import (
	"github.com/dgrijalva/jwt-go"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Id string `json:"id"`

	Username           uint   `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Picture            bool   `json:"picture"`
	jwt.StandardClaims `swaggerignore:"true"`
}
