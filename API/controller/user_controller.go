package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ksw95/GoIndustrialProject/API/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

//New User Account Creation
func Register(c echo.Context) {
	//Receive parameters from client
	user := models.Account{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return newResponse(c, "Internal Error", "false", http.StatusInternalServerError, nil)
	}
	//Check username whether exist in database
	//Create new entry in database
	
}

//User Login
func Login(c echo.Context) {
	//Receive params from client
	//Check username and password
	//Assign JWT with UserCond info
	fmt.Println("This is for login") //delete after changing the function
}

//Updating User Conditions
func UserCondition() {
	//Receive params from client
	//Check Account via JWT
	//Update into database
	fmt.Println("This is for user condition") //delete after changing the function
}
