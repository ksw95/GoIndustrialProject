
// start server
e := echo.New()

e.GET("/api/v0/check", controller.PwCheck(c) 
})

s := http.Server{Addr: ":" + port, Handler: e}

err := s.ListenAndServe()

/// get params
func GetAllListingIndex(c echo.Context) error {
	// can only return listing results, commentUser and commentItem
	itemName := c.QueryParam("name")
	filterUsername := c.QueryParam("filter")
	fmt.Println(itemName, filterUsername)
	return c.JSON(httpStatus, nil)
}

/// get params
// return json
// note the :id
e.PUT("/api/v0/db/completed/:id", func(c echo.Context) error {
	return controller.Completed(c, &dbHandler1)
})

func GetAllListingIndex(c echo.Context) error {
	// can only return listing results, commentUser and commentItem
	itemID := c.Param("id")
	fmt.Println(itemName, filterUsername)
	responseJson := model.DataPacketSimple{
		msg,
		resBool,
	}
	return c.JSON(httpStatus, responseJson)
}


/// encoding url params
q := make(url.Values)
q.Set("id", c.Param("id"))
q.Set("db", "ItemListing")
dataPacket1, err1 := TapApi(http.MethodGet, "", "db/info?id="+c.Param("id")+"&db=ItemListing", sessionMgr)
