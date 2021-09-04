package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ksw95/GoIndustrialProject/API/controller"

	_ "github.com/go-sql-driver/mysql" // go mod init api_server.go
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() (http.Server, *echo.Echo, *controller.DBHandler, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	err := godotenv.Load("go.env")
	if err != nil {
		fmt.Println(err)
	}

	dbHandler := controller.OpenDB(os.Getenv("DATABASE"))
	dbHandler.ApiKey = os.Getenv("API_KEY")

	e := echo.New()

	// not sure if it will work, need to double check
	// middleware to check apikey and server readiness and reply if not ready or key is wrong
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !dbHandler.ReadyForTraffic {
				fmt.Println("API is accessed, but sql server is unavailable")
				responseJson := struct {
					Msg     string //message
					ResBool string //boolean response
				}{
					"not ready for traffic",
					"false",
				}
				// encode to json and send
				return c.JSON(503, responseJson)
			}
			if c.QueryParam("key") != dbHandler.ApiKey { ////
				// encode to json and send
				return echo.NewHTTPError(http.StatusUnauthorized, "")

			}
			return next(c)
		}
	})

	//check if mysql is active
	go func(dbHandler *controller.DBHandler) {
		for {
			results, err1 := dbHandler.DB.Query("Select * FROM Restaurant WHERE ID = ?", 1)
			if err1 != nil {
				dbHandler.ReadyForTraffic = false
				fmt.Println("unable to contact mysql server, error:", err1.Error())
			} else {
				dbHandler.ReadyForTraffic = true
				results.Close()
			}
			time.Sleep(120 * time.Second)
		}
	}(dbHandler)

	// other echo middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//handler functions
	e.GET("/api/v0/health", controller.HealthCheckLiveness)

	e.GET("/api/v0/GetRestaurant", dbHandler.GetRestaurant)

	e.GET("/api/v0/GetRestaurantAll", dbHandler.GetRestaurantAll)

	e.GET("/api/v0/InsertRestaurant", dbHandler.InsertRestaurant)

	e.GET("/api/v0/SearchRestaurant", dbHandler.SearchRestaurant)

	e.GET("/api/v0/GetFoodShopID", dbHandler.GetFoodShopID)

	e.GET("/api/v0/InsertFood", dbHandler.InsertFood)

	e.GET("/api/v0/EditFood", dbHandler.EditFood)

	//server
	port := os.Getenv("PORT")
	fmt.Println("listening at port " + port)
	s := http.Server{Addr: ":" + port, Handler: e}

	return s, e, dbHandler, nil
}

func main() {
	s, e, dbHandler, _ := StartServer()
	defer dbHandler.DB.Close()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
