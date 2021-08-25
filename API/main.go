package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/Mechwarrior1/PGL_backend/controller"
	"github.com/Mechwarrior1/PGL_backend/model"
	"github.com/Mechwarrior1/PGL_backend/mysql"

	_ "github.com/go-sql-driver/mysql" // go mod init api_server.go
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() (http.Server, *echo.Echo, *model.DBHandler, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	err := godotenv.Load("go.env")
	if err != nil {
		fmt.Println(err)
	}

	dbHandler := controller.DBHandler{
		mysql.OpenDB(),
		os.Getenv("API_KEY"),
		false,
	}

	e := echo.New()

	// not sure if it will work, need to double check
	// middleware to check apikey and server readiness and reply if not ready or key is wrong
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !dbHandler.ReadyForTraffic {
				fmt.Println("API is accessed, but is unable to contact sql server")
				responseJson := controller.DataPacketSimple{
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

	// other echo middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//handler functions
	e.GET("/api/v0/GetRestaurant", dbHandler.GetRestaurant)

	e.GET("/api/v0/GetRestaurantAll", dbHandler.GetRestaurantAll)

	e.GET("/api/v0/SearchRestaurant", dbHandler.SearchRestaurant)

	e.GET("/api/v0/GetFoodShopID", dbHandler.GetFoodShopID)

	e.GET("/api/v0/health", controller.HealthCheckLiveness)

	//server
	port := os.Getenv("PORT")
	fmt.Println("listening at port " + port)
	s := http.Server{Addr: ":" + port, Handler: e}

	return s, e, &dbHandler, nil
}

func main() {
	s, e, dbHandler, _ := StartServer()
	defer dbHandler.DBController.DB.Close()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
