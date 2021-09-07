// this file manages the sessions and user information.
package session

import (
	"net/http"
	"time"

	"github.com/ksw95/GoIndustrialProject/Client/models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type (
	// a struct to manage sessions
	Session struct {
		MapSession *map[string]*SessionStruct // maps UUID to user session
		Client     ClientDo
		ApiKey     string //get key from env
		BaseURL    string //get url from env
	}

	SessionStruct struct {
		LastActive int64
		UserCon    *models.UserCond
		Success    string
		SuccessMsg string
	}

	// logger struct {
	// 	c1 chan string
	// 	c2 chan string
	// }

	ClientDo interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

//new session for users
func (s *Session) NewSession(c echo.Context, userCond *models.UserCond) *SessionStruct {
	newUuid := uuid.NewV4().String()

	//log uuid session to session map
	newSession := &SessionStruct{time.Now().Unix(), userCond, "", ""}
	(*s.MapSession)[newUuid] = newSession

	// store uuid in cookie
	NewCookie(c, 3, newUuid)

	return newSession
}

//new session for users
func (s *Session) NewEmptySession(c echo.Context) *SessionStruct {

	// var userCond models.UserCond

	// dummy userCond
	userCond := models.UserCond{
		"new user",
		"7/9/2021",
		4000,
		"Diabetic",
		"Halal",
		"Vegan",
	}

	return s.NewSession(c, &userCond)
}

// logs the session in the sessionmanager.
func (s *Session) CheckSession(c echo.Context) *SessionStruct {
	cookieVal, err := c.Cookie("foodiepandaCookie")

	// cookie not found, new session created
	if err != nil {
		return s.NewEmptySession(c)
	}

	sessionStruct, ok1 := (*s.MapSession)[cookieVal.Value] // check if previous session is around

	// session not found, new session created
	if !ok1 || cookieVal.Value == "" {
		return s.NewEmptySession(c)
	}

	return sessionStruct
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
