package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"text/template"

	"github.com/ksw95/GoIndustrialProject/Client/session"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	// GetDoFunc fetches the mock client's `Do` func
	Client HTTPClient
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type (
	// Custom type that allows setting the func that our Mock Do func will run instead
	MockDoFunc func(req *http.Request) (*http.Response, error)

	// MockClient is the mock client
	MockClient struct {
		MockDo MockDoFunc
	}

	Template struct {
		templates *template.Template
	}

	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

// provides the wrapper required by handlers
func getDependency(address string, encode io.Reader) (*session.Session, *httptest.ResponseRecorder, *http.Request, *session.SessionStruct, *echo.Echo, echo.Context) {

	mapSession := make(map[string]*session.SessionStruct)
	sessionMgr := &session.Session{
		MapSession: &mapSession,
		ApiKey:     "key",
		Client:     &http.Client{},
	}

	//mock form values to function post
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, address, encode)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	e := echo.New()
	c := e.NewContext(req, rec)
	sessionStruct, uuid := sessionMgr.NewEmptySession(c)

	req.AddCookie(&http.Cookie{
		Name:   "foodiepandaCookie",
		Value:  uuid,
		MaxAge: 300,
	})

	c = e.NewContext(req, rec)

	tem := &Template{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}
	e.Renderer = tem

	return sessionMgr, rec, req, sessionStruct, e, c
}

func TestIndex_GET(t *testing.T) {

	client := &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 202,
			}, nil
		},
	}

	// get dependencies
	sessionMgr, rec, _, _, _, c := getDependency("/", nil)
	sessionMgr.Client = client

	// fmt.Println(rec.Body)

	if assert.NoError(t, Index_GET(c, sessionMgr)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestIndex_POST(t *testing.T) {

	client := &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 202,
			}, nil
		},
	}

	//insert form value
	f := make(url.Values)
	f.Set("cat", "Food")
	f.Set("query", "item")

	// get dependencies
	sessionMgr, rec, _, _, _, c := getDependency("/", strings.NewReader(f.Encode()))
	sessionMgr.Client = client

	// fmt.Println(rec.Body)

	if assert.NoError(t, Index_POST(c, sessionMgr)) {
		assert.Equal(t, http.StatusSeeOther, rec.Code)
		fmt.Println(c.Request().URL)
	}
}

func TestSearchPage_GET_Food(t *testing.T) {
	// create dependencies

	// mock client Do for handler
	// build our response JSON

	// create a new reader with that JSON

	client := &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {

			querySearch := req.URL.Query().Get("search")
			queryType := req.URL.Query().Get("type")

			assert.Equal(t, "item", querySearch)
			assert.Equal(t, "Food", queryType)

			dataResponse := []interface{}{getDummyData("Food", 1),
				getDummyData("Food", 2),
				getDummyData("Food", 3),
			}

			searchResult := struct {
				Msg     string
				ResBool string
				Data    []interface{}
			}{
				"ok",
				"true",
				dataResponse,
			}

			jsonResponse, _ := json.Marshal(searchResult)
			r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

			return &http.Response{
				StatusCode: 202,
				Body:       r,
			}, nil
		},
	}

	//insert form value
	q := make(url.Values)
	q.Set("cat", "Food")
	q.Set("query", "item")

	sessionMgr, rec, _, _, _, c := getDependency("/search?"+q.Encode(), nil)
	sessionMgr.Client = client

	// fmt.Println(rec.Body)

	if assert.NoError(t, SearchPage_GET(c, sessionMgr)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(c.Request().URL)
	}
}

func TestSearchPage_GET_Restaurant(t *testing.T) {
	// create dependencies

	// mock client Do for handler
	// build our response JSON

	// create a new reader with that JSON

	client := &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {

			querySearch := req.URL.Query().Get("search")
			queryType := req.URL.Query().Get("type")

			assert.Equal(t, "item", querySearch)
			assert.Equal(t, "Restaurant", queryType)

			dataResponse := []interface{}{getDummyData("Restaurant", 1),
				getDummyData("Restaurant", 2),
				getDummyData("Restaurant", 3),
			}

			searchResult := struct {
				Msg     string
				ResBool string
				Data    []interface{}
			}{
				"ok",
				"true",
				dataResponse,
			}

			jsonResponse, _ := json.Marshal(searchResult)
			r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

			return &http.Response{
				StatusCode: 202,
				Body:       r,
			}, nil
		},
	}

	//insert form value
	q := make(url.Values)
	q.Set("cat", "Restaurant")
	q.Set("query", "item")

	sessionMgr, rec, _, _, _, c := getDependency("/search?"+q.Encode(), nil)
	sessionMgr.Client = client

	// fmt.Println(rec.Body)

	if assert.NoError(t, SearchPage_GET(c, sessionMgr)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(c.Request().URL)
	}
}

func getDummyData(db string, id int) (newMap map[string]interface{}) {
	newMap = make(map[string]interface{})
	switch db {
	case "UserCond":
		newMap["Username"] = id
		newMap["LastLogin"] = "20-7-2021"
		newMap["MaxCalories"] = 2500
		newMap["Diabetic"] = "Not Diabetic"
		newMap["Halal"] = "Not Halal"
		newMap["Vegan"] = "Not Vegan"
	case "Food":
		newMap["ID"] = id
		newMap["Name"] = "Name"
		newMap["ShopID"] = 1
		newMap["Calories"] = 800
		newMap["Description"] = "Description"
		newMap["Sugary"] = "Not Sugary"
		newMap["Halal"] = "Not Halal"
		newMap["Vegan"] = "Not Vegan"
	case "History":
		newMap["ID"] = id
		newMap["Username"] = "username"
		newMap["FoodPurchase"] = "20-7-2021"
		newMap["Price"] = 3.7
		newMap["DeliveryMode"] = "Walking"
		newMap["Distance"] = 2.2
		newMap["CaloriesBurned"] = 120
	case "Account":
		newMap["Username"] = "Username"
		newMap["Password"] = "Password"
	}
	return
}
