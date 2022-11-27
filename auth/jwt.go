package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JWTSecret = []byte("secret")

func GenerateJWT(id uint, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["Email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func JWT() echo.MiddlewareFunc {
	c := middleware.DefaultJWTConfig
	c.SigningKey = JWTSecret
	return middleware.JWTWithConfig(c)
}
