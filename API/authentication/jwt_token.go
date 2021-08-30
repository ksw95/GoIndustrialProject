package authentication

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(c echo.Context, username string) {
	claims := jwt.MapClaims{
		"username": username,
		// UserCond data
		"expiry": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("GoIndustrialProject"))
	if err != nil {
		log.Fatalf("token creation error %s\n", err)
	}
	cookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}
