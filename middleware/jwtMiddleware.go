package middleware

import (
	"Chiprek/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SCREAT_JWT),
})

// Create Token Admin
func CreateToken(adminId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin_id"] = adminId
	claims["username"] = "admin"
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}

// Create Token User
func CreateTokenUser(CustomerId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["customer_id"] = CustomerId
	claims["name"] = name
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}

// func IsAdmin(c echo.Context) (int, error) {
// 	user := c.Get("user").(*jwt.Token)
// 	if !user.Valid {
// 		return 0, echo.NewHTTPError(401, "Unauthorized")
// 	}
// 	claims := user.Claims.(jwt.MapClaims)
// 	if claims["username"] != "admin" {
// 		return 0, echo.NewHTTPError(401, "Unauthorized")
// 	}

// 	adminId := int(claims["admin_id"].(float64))

// 	return adminId, nil
// }

func IsCustomer(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["name"] == "admin" {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	customerId := int(claims["customer_id"].(float64))

	return customerId, nil
}
