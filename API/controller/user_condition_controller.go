package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ksw95/GoIndustrialProject/API/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func (dbHandler *DBHandler) Insert(c echo.Context) error {
        userC := models.UserCond{}
	err := json.NewDecoder(c.Request().Body).Decode(&userC)
	if err != nil {
		fmt.Println(err.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	// prepare statement to insert record
	tx, err := controller.DBHandler.DB.Begin()
	if err != nil {
		return err
	}

	//first statement
	stmt, err1 := tx.Prepare("INSERT INTO Condition VALUES (?, DATE_ADD(NOW(), INTERVAL 8 HOUR), ?, ?, ?, ?))")

	if err1 == nil {
		fmt.Println(err1)
	}

	_, err = stmt.Exec(userC.Username, userC.MaxCalories, userC.Diabetic, userC.Halal, userC.Vegan)

	stmt.Close()

	switch err {
	case nil:
		_ = tx.Commit()
		return newResponse(c, "ok", "true", http.StatusOK, nil)
	default:
		tx.Rollback()
		return newResponse(c, "rolled back", "false", http.StatusBadRequest, nil)
	}
}

func (dbHandler *DBHandler) Update(c echo.Context) error {
        userC := models.UserCond{}
	err := json.NewDecoder(c.Request().Body).Decode(&userC)
	if err != nil {
		fmt.Println(err.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	// prepare statement to insert record
	tx, err := controller.DBHandler.DB.Begin()
	if err != nil {
		return err
	}

	//first statement
	stmt, err1 := tx.Prepare("UPDATE Condition " +
		"SET =DATE_ADD(NOW(), INTERVAL 8 HOUR), MaxCalories=?, Diabetic=?, Halal=?, Vegan=? " +
		"WHERE Username=?")

	if err1 == nil {
		_, err = stmt.Exec(userC.MaxCalories, userC.Diabetic, userC.Halal, userC.Vegan, userC.Username)
	}

	stmt.Close()

	switch err {
	case nil:
		_ = tx.Commit()
		return newResponse(c, "ok", "true", http.StatusOK, nil)
	default:
		tx.Rollback()
		return newResponse(c, "rolled back", "false", http.StatusBadRequest, nil)
	}
}

func (dbHandler *DBHandler) Get(c echo.Context) error {
        userC := models.UserCond{}
	//get id param
	username := c.QueryParam("Username")
	if username == "" {
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	// query mysql
	results, err1 := controller.DBHandler.DB.Query("SELECT * FROM MemberType WHERE Username=?", username)

	if err1 != nil {
		fmt.Println(err1.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}
	defer results.Close()

	//scan mysql result
	results.Next()
	err2 := results.Scan(&userC.Username, &userC.MaxCalories, &userC.Diabetic, &userC.Halal, &userC.Vegan)

	if err2 != nil {
		fmt.Println(err2.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	//return json
	return newResponse(c, "ok", "true", http.StatusOK, &[]interface{}{restaurant})
}
