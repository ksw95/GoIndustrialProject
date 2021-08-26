package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

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
	// load variables and mock database
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
