package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksw95/GoIndustrialProject/Client/models"
	"github.com/ksw95/GoIndustrialProject/Client/session"
	"github.com/labstack/echo"
)

// the function to access the rest api
// requires the method and datapacket
// returns any courseinfo and error
func TapApi(httpMethod string, jsonData interface{}, url string, sessionMgr *session.Session) (*[]byte, error) {
	//complete url with apikey
	url = sessionMgr.BaseURL + url + "&key=" + sessionMgr.ApiKey

	var request *http.Request

	// prepare request depending on if there is json data to be sent
	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		jsonValueMarshal := bytes.NewBuffer(jsonValue)
		request, _ = http.NewRequest(httpMethod, url, jsonValueMarshal)

	} else {
		request, _ = http.NewRequest(httpMethod, url, nil)
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := sessionMgr.Client.Do(request)

	if err != nil {
		fmt.Println("TapApi failed with error:", err.Error())
		return nil, errors.New("http request failed with " + err.Error())
	}

	data, err1 := ioutil.ReadAll(response.Body) //
	response.Body.Close()

	return &data, err1
}

// handler function, for the index page
func Index_GET(c echo.Context, sessionMgr *session.Session) error {
	userSes := sessionMgr.CheckSession(c)

	// session.Update

	return c.Render(http.StatusOK, "index.gohtml", userSes)
}

// handler function, for the index page
// when posting, takes form params and redirect to search page
func Index_POST(c echo.Context, sessionMgr *session.Session) error {
	form, _ := c.FormParams()

	postSearch := form["search"][0]
	postCat := form["cat"][0] //food or restaurant
	url := "/view?search=" + postSearch + "&cat=" + postCat
	return c.Redirect(http.StatusSeeOther, url)

}

// handler function, for the Search page
//
func SearchPage_GET(c echo.Context, sessionMgr *session.Session) error {
	userSes := sessionMgr.CheckSession(c)

	form, _ := c.FormParams()
	postSearch := form["search"][0]
	postCat := form["cat"][0] //food or restaurant category

	url := "search=" + postSearch + "&type=" + postCat

	dataByte, err := TapApi(http.MethodGet, "", "v0/SearchRestaurant?"+url, sessionMgr)
	if err != nil {
		fmt.Println("TapApi failed with error:", err.Error())
		return errors.New("http request failed with " + err.Error())
	}

	if postCat == "food" {
		var searchResult []models.Food
		json.Unmarshal(*dataByte, &searchResult)

		dataInsert := struct {
			UserData       models.UserCond
			SearchResult   []models.Food
			PaginationBool bool
			FoodBool       bool
		}{
			*(userSes.UserCon),
			searchResult,
			false,
			true,
		}

		return c.Render(http.StatusOK, "searchpage.gohtml", dataInsert)
	}

	var searchResult []models.Restaurant
	json.Unmarshal(*dataByte, &searchResult)

	dataInsert := struct {
		UserData       models.UserCond
		SearchResult   []models.Restaurant
		PaginationBool bool
		FoodBool       bool
	}{
		*(userSes.UserCon),
		searchResult,
		false,
		false,
	}

	return c.Render(http.StatusOK, "searchpage.gohtml", dataInsert)

}
