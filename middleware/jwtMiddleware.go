package middleware

import (
	"Chiprek/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	mid "github.com/labstack/echo/middleware"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SCREAT_JWT),
})

func CreateToken(adminId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin_id"] = adminId
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}
