// this file manages the sessions and user information.
package session

import (
	"fmt"
	"net/http"

	"github.com/Mechwarrior1/PGL_frontend/jwtsession"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type (
	// a struct to manage sessions
	Session struct {
		MapSession *map[string]SessionStruct // maps UUID to user session
		ApiKey     string
		Client     ClientDo
		BaseURL    string
	}

	SessionStruct struct {
		LastActive int64
		UserC      *UserCond
		Success	string
		SuccessMsg	string
	}

	UserCond struct {
		Username    string
		MaxCalories int
		Diabetic    bool
		Halal       bool
		Vegan       bool
	}

	// logger struct {
	// 	c1 chan string
	// 	c2 chan string
	// }

	ClientDo interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

// logs the session in the sessionmanager.
func (s *Session) CheckSession(c echo.Context) SessionStruct {
	cookieVal, err := c.Cookie("foodiepandaCookie")

	sessionStruct, ok := (*s.MapSession)[cookieVal] // check if previous session is around

	if !ok || cookieVal = ""{ 
		newUuid := uuid.NewV4().String()

		//new session id for users not logged in
		(*s.MapSession)[newUuid] = SessionStruct{time.Now().Unix(), nil}
		NewCookie(c, 3, newJwt)

		return claims // return new claims for user, since old session got terminated
	}

	return jwtClaim
}

// function deletes the session, based on the session id string.
func (s *Session) DeleteSession(uuid string) {
	delete(*s.MapSession, uuid)
}

// a function to generate a new cookie, with the session id as cookie value.
func NewCookie(c echo.Context, expiry int, id string) { //make a new cookie.
	foodiepandaCookie := &http.Cookie{
		Name:   "foodiepandaCookie",
		Value:  id,
		MaxAge: expiry,
		Path:   "/",
	}
	c.SetCookie(foodiepandaCookie)
}

func ExpCookie(c echo.Context) { //make a new cookie.
	NewCookie(c, -1, "")
}
