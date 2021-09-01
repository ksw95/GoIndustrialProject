package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// the function to access the rest api
// requires the method and datapacket
// returns any courseinfo and error
func TapApi(httpMethod string, jsonData interface{}, url string, sessionMgr *session.Session) (*map[string]interface{}, error) {
	url = sessionMgr.BaseURL + url
	var request *http.Request
	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		jsonValueMarshal := bytes.NewBuffer(jsonValue)
		request, _ = http.NewRequest(httpMethod, url, jsonValueMarshal)
	} else {
		request, _ = http.NewRequest(httpMethod, url, nil)
	}

	request.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	response, err := sessionMgr.Client.Do(request)
	mapInterface := make(map[string]interface{})
	if err != nil {
		fmt.Println("TapApi failed with error:", err.Error())
		return &mapInterface, errors.New("http request failed with " + err.Error())

	} else {

		data1, err := ioutil.ReadAll(response.Body) //

		if err != nil {
			return &mapInterface, errors.New("ioutil failed to read, error: " + err.Error())
		}

		json.Unmarshal(data1, &mapInterface)
		response.Body.Close()

		if mapInterface["ErrorMsg"] != "nil" {
			return &mapInterface, errors.New("error")
		}

		// if data, ok := mapInterface["DataInfo"]; ok && len(data.([]interface{}) == 1{
		// 	dataTemp := data.([]interface{})[0].(map[string]interface{})
		// 	mapInterface["DataInfo"] = dataTemp
		// }

		return &mapInterface, nil
	}
}

// handler function, for the index page
func Index_GET(c echo.Context, jwtWrapper *jwtsession.JwtWrapper, sessionMgr *session.Session) error {
	userSes := CheckSession(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	// session.Update
	return c.Render(http.StatusOK, "index.gohtml", userSes)
}

// handler function, for the index page
// when posting, takes form params and redirect to search page
func Index_POST(c echo.Context, jwtWrapper *jwtsession.JwtWrapper, sessionMgr *session.Session) error {
	form, _ := c.FormParams()

	postSearch := form["search"][0]
	postCat := form["cat"][0] //food or restaurant
	url := "/view?search=" + postSearch + "&cat=" + postCat
	return c.Redirect(http.StatusSeeOther, url)

}
