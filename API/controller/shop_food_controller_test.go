package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ksw95/GoIndustrialProject/API/models"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql" // go mod init api_server.go
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*DBHandler, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbHandler := &DBHandler{
		db,
		"",
		true}

	return dbHandler, mock
}

func TestGetRestaurant(t *testing.T) {

	// load dependencies
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock DB
	// bPassword, _ := bcrypt.GenerateFromPassword([]byte("john"), bcrypt.MinCost)

	rows := sqlmock.NewRows([]string{"ID", "Name", "Description", "Address", "PostalCode"}).
		AddRow(1, "curry", "good curry", "blk 123", 123)
	query := "Select \\* FROM Restaurant WHERE ID = \\?"
	mock.ExpectQuery(query).WillReturnRows(rows)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("id", "1")
	fmt.Println("")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetRestaurant?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetRestaurant(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "true", json_map["ResBool"])

		json_map_data := json_map["Data"].([]interface{})[0].(map[string]interface{})
		assert.Equal(t, float64(1), json_map_data["ID"].(float64))
		assert.Equal(t, "curry", json_map_data["Name"].(string))

	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRestaurant_Fail(t *testing.T) {
	// load dependencies
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock DB
	// bPassword, _ := bcrypt.GenerateFromPassword([]byte("john"), bcrypt.MinCost)

	query := "Select \\* FROM Restaurant WHERE ID = \\?"
	mock.ExpectQuery(query)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "")
	q.Set("id", "1")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetRestaurant?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "Entry not found", json_map["Msg"])
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRestaurant_Fail2(t *testing.T) {
	// load dependencies
	dbHandler, _ := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetRestaurant?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "id field is empty", json_map["Msg"])
	}

}

func TestGetRestaurantAll(t *testing.T) {
	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	//sql mock
	rows := sqlmock.NewRows([]string{"ID", "Name", "Description", "Address", "PostalCode"}).
		AddRow(1, "curry", "good curry", "blk 1", 1).
		AddRow(2, "curry2", "good curry2", "blk 2", 2).
		AddRow(3, "curry3", "good curry3", "blk 3", 3)
	query := "Select \\* FROM Restaurant"
	mock.ExpectQuery(query).WillReturnRows(rows)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetRestaurantAll?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetRestaurantAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "true", json_map["ResBool"])

		json_map_data1 := json_map["Data"].([]interface{})[0].(map[string]interface{})
		assert.Equal(t, float64(1), json_map_data1["ID"].(float64))
		assert.Equal(t, "curry", json_map_data1["Name"].(string))

		json_map_data2 := json_map["Data"].([]interface{})[1].(map[string]interface{})
		assert.Equal(t, float64(2), json_map_data2["ID"].(float64))
		assert.Equal(t, "curry2", json_map_data2["Name"].(string))

		json_map_data3 := json_map["Data"].([]interface{})[2].(map[string]interface{})
		assert.Equal(t, float64(3), json_map_data3["ID"].(float64))
		assert.Equal(t, "curry3", json_map_data3["Name"].(string))

	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRestaurantAll_fail(t *testing.T) {
	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	//sql mock, intentionally make an error
	// rows := sqlmock.NewRows([]string{"ID", "Name", "Description", "Address", "PostalCode"})
	query := "Select \\* FROM Restaurant" //
	mock.ExpectQuery(query)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetRestaurantAll?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetRestaurantAll(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "Bad Request", json_map["Msg"])

	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertRestaurant(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying max id
	query := "SELECT MAX\\(ID\\) FROM Restaurant" //for MaxID query
	rows := sqlmock.NewRows([]string{"ID"}).
		AddRow(1)
	mock.ExpectQuery(query).WillReturnRows(rows)

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "INSERT INTO Restaurant VALUES \\(\\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(1, "curry", "123", "456", 123).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	// json payload to api
	restuarant := models.Restaurant{
		1,
		"curry",
		"123",
		"456",
		123,
	}

	payloadJson, _ := json.Marshal(restuarant)

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/InsertRestaurant", strings.NewReader(string(payloadJson)))
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.InsertRestaurant(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "true")

	}

}

func TestInsertRestaurant_Fail(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying max id
	query := "SELECT MAX\\(ID\\) FROM Restaurant" //for MaxID query
	rows := sqlmock.NewRows([]string{"ID"}).
		AddRow(1)
	mock.ExpectQuery(query).WillReturnRows(rows)

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "INSERT INTO Restaurant VALUES \\(\\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec()

	mock.ExpectRollback()

	// json payload to api

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/InsertRestaurant", nil)
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.InsertRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")

	}
}

func TestEditRestaurant(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "UPDATE Restaurant SET Name=\\?, Description=\\?, Address=\\?, PostalCode=\\? WHERE ID=\\?"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs("curry", "123", "456", 123, 1).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	// json payload to api
	restuarant := models.Restaurant{
		1,
		"curry",
		"123",
		"456",
		123,
	}

	payloadJson, _ := json.Marshal(restuarant)

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/EditRestaurant", strings.NewReader(string(payloadJson)))
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.EditRestaurant(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "true")

	}

}

func TestEditRestaurant_Fail(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "UPDATE Restaurant SET Name=\\?, Description=\\?, Address=\\?, PostalCode=\\? WHERE ID=\\?"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec()

	mock.ExpectRollback()

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/EditRestaurant", nil)
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.EditRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")

	}

}

func TestSearchRestaurant(t *testing.T) {

	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying, sql payload
	rows := sqlmock.NewRows([]string{"ID", "Name", "Description", "Address", "PostalCode"}).
		AddRow(1, "curry", "curry", "1", 123).
		AddRow(2, "curry", "2", "2", 123).
		AddRow(3, "3", "curry", "3", 123)
	query := "Select \\* FROM Restaurant"
	mock.ExpectQuery(query).WillReturnRows(rows)

	//url payload
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("search", "curry")
	q.Set("type", "restaurant")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/SearchRestaurant?"+q.Encode(), nil)
	c := e.NewContext(req, rec)

	//inject dependencies and test
	if assert.NoError(t, dbHandler.SearchRestaurant(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		json_map2 := json_map["Data"].([]interface{})

		assert.Equal(t, json_map2[0].(map[string]interface{})["ID"].(float64), float64(1))
		assert.Equal(t, json_map2[1].(map[string]interface{})["ID"].(float64), float64(2))
		assert.Equal(t, json_map2[2].(map[string]interface{})["ID"].(float64), float64(3))
	}
}

func TestSearchRestaurant_Fail(t *testing.T) {

	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying, sql payload
	query := "Select \\* FROM Restaurant"
	mock.ExpectQuery(query)

	//url payload
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("search", "curry")
	q.Set("type", "restaurant")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/SearchRestaurant?"+q.Encode(), nil)
	c := e.NewContext(req, rec)

	//inject dependencies and test
	if assert.NoError(t, dbHandler.SearchRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")

	}
}

func TestSearchFood(t *testing.T) {

	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying, sql payload
	rows := sqlmock.NewRows([]string{"ID", "Name", "ShopID", "Calories", "Description", "Sugary", "Halal", "Vegan"}).
		AddRow(1, "curry", 1, 1, "curry", "f", "f", "f").
		AddRow(2, "12345", 1, 1, "curry", "f", "f", "f").
		AddRow(3, "curry", 1, 1, "12345", "f", "f", "f")
	query := "Select \\* FROM Food"
	mock.ExpectQuery(query).WillReturnRows(rows)

	//url payload
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("search", "curry")
	q.Set("type", "food")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/SearchRestaurant?"+q.Encode(), nil)
	c := e.NewContext(req, rec)

	//inject dependencies and test
	if assert.NoError(t, dbHandler.SearchRestaurant(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		json_map2 := json_map["Data"].([]interface{})

		assert.Equal(t, json_map2[0].(map[string]interface{})["ID"].(float64), float64(1))
		assert.Equal(t, json_map2[1].(map[string]interface{})["ID"].(float64), float64(3))
		assert.Equal(t, json_map2[2].(map[string]interface{})["ID"].(float64), float64(2))

	}
}

func TestSearchFood_Fail(t *testing.T) {

	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying, sql payload
	query := "Select \\* FROM Food"
	mock.ExpectQuery(query)

	//url payload
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("search", "curry")
	q.Set("type", "food")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/SearchRestaurant?"+q.Encode(), nil)
	c := e.NewContext(req, rec)

	//inject dependencies and test
	if assert.NoError(t, dbHandler.SearchRestaurant(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")

	}
}

func TestGetFoodShopID(t *testing.T) {
	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock DB
	// bPassword, _ := bcrypt.GenerateFromPassword([]byte("john"), bcrypt.MinCost)

	rows := sqlmock.NewRows([]string{"ID", "Name", "ShopID", "Calories", "Description", "Sugary", "Halal", "Vegan"}).
		AddRow(1, "curry", 1, 1, "curry", "f", "f", "f")
	query := "Select \\* FROM Food WHERE ShopID = \\?"
	mock.ExpectQuery(query).WillReturnRows(rows)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("id", "1")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetFoodShopID?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetFoodShopID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "true", json_map["ResBool"])

		json_map_data := json_map["Data"].([]interface{})[0].(map[string]interface{})
		assert.Equal(t, float64(1), json_map_data["ID"].(float64))
		assert.Equal(t, "curry", json_map_data["Name"].(string))

	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFoodShopID_Fail(t *testing.T) {
	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock DB
	query := "Select \\* FROM Food WHERE ShopID = \\?"
	mock.ExpectQuery(query)

	// making the call to api and encode variables
	q := make(url.Values)
	q.Set("key", "123")
	q.Set("id", "1")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v0/GetFoodShopID?"+q.Encode(), nil)

	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.GetFoodShopID(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)
		assert.Equal(t, "false", json_map["ResBool"])

	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertFood(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying max id
	query := "SELECT MAX\\(ID\\) FROM Food" //for MaxID query
	rows := sqlmock.NewRows([]string{"ID"}).
		AddRow(1)
	mock.ExpectQuery(query).WillReturnRows(rows)

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "INSERT INTO Food VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(1, "curry", 1, 1, "123", "123", "123", "123").WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	// json payload to api
	restuarant := models.Food{
		1,
		"curry",
		1,
		1,
		"123",
		"123",
		"123",
		"123",
	}

	payloadJson, _ := json.Marshal(restuarant)

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/InsertFood", strings.NewReader(string(payloadJson)))
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.InsertFood(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "true")

	}

}

func TestInsertFood_Fail(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock for querying max id
	query := "SELECT MAX\\(ID\\) FROM Food" //for MaxID query
	mock.ExpectQuery(query)

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "INSERT INTO Food VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec()

	mock.ExpectRollback()

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/InsertFood", nil)
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.InsertFood(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")
	}
}

func TestEditFood(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "UPDATE Food SET Name=\\?, ShopID=\\?, Calories=\\?, Description=\\?, Sugary=\\?, Halal=\\?, Vegan=\\? WHERE ID=\\?"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs("curry", 1, 1, "123", "123", "123", "123", 1).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	// json payload to api
	restuarant := models.Food{
		1,
		"curry",
		1,
		1,
		"123",
		"123",
		"123",
		"123",
	}

	payloadJson, _ := json.Marshal(restuarant)

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/EditFood?", strings.NewReader(string(payloadJson)))
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.EditFood(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "true")

	}

}

func TestEditFood_Fail(t *testing.T) {

	// load variables and mock database
	dbHandler, mock := NewMock()
	defer dbHandler.DB.Close()

	e := echo.New()

	// mock sql for inserting value
	mock.ExpectBegin()
	query2 := "UPDATE Food SET Name=\\?, ShopID=\\?, Calories=\\?, Description=\\?, Sugary=\\?, Halal=\\?, Vegan=\\? WHERE ID=\\?"

	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec()
	mock.ExpectRollback()

	// test api
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v0/EditFood?", nil)
	c := e.NewContext(req, rec)

	if assert.NoError(t, dbHandler.EditFood(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		//check response
		json_map := make(map[string]interface{})
		json.NewDecoder(rec.Body).Decode(&json_map)

		assert.Equal(t, json_map["ResBool"], "false")

	}

}

func TestMerge(t *testing.T) {
	_, result, err := MergeSort([]int{5, 0, 3, 8, 20, 14, 15, 9, 2, 11}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for idx, i := range []int{1, 8, 2, 0, 3, 7, 9, 5, 6, 4} {
		assert.Equal(t, result[idx], i)
		assert.Equal(t, err, nil)
	}

}

func TestMerge_Fail(t *testing.T) {
	_, _, err := MergeSort([]int{5, 0, 3, 8, 20, 14, 15, 9, 2, 11}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	assert.Equal(t, err, errors.New("array not of same length"))

}
