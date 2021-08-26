package controller

import "fmt"

//New User Account Creation
func Register() {
	//Receive parameters from client
	//Check username whether exist in database
	//Create new entry in database
	fmt.Println("This is for registering") //delete after changing the function
}

//User Login
func Login() {
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
