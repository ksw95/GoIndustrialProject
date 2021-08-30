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
func (dbHandler *DBHandler) Register(c echo.Context) {
	//Receive parameters from client
	user := models.Account{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
		return
	}
	hashPW, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		newResponse(c, "Internal Error", "false", http.StatusInternalServerError, nil)
		return
	}
	user.Password = string(hashPW)
	//Check username whether exist in database
	//Create new entry in database
	_, err = dbHandler.DB.Exec("INSERT INTO my_db.Account VALUES (?,?)", user.Username, user.Password)
	if err != nil {
		newResponse(c, "Failed to create new account", "false", http.StatusForbidden, nil)
		return
	}
	newResponse(c, "Sucessfully created", "true", http.StatusCreated, nil)
	return
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
