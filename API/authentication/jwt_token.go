package authentication

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(c echo.Context, username string) string {
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
	return tokenString
}

func ExtractToken(c echo.Context) string {
	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		return err
	}
	return cookie.Value
}

func ExtractTokenID(c echo.Context) (uint32, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("GoIndustrialProject"), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["username"]
		if err != nil {
			return 0, err
		}

		fmt.Println("jwt-token")
		return username, nil
	}

	return 0, nil
}
