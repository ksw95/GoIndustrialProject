package models

type UserCond struct {
	Username    string
	MaxCalories int
	Diabetic    bool
	Halal       bool
	Vegan       bool
}

/*
func (userC *UserCond) Insert(w http.ResponseWriter, r *http.Request) {
	// Check valid key
	if !validKey(r) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Invalid key"))
		return
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			// Read the string sent to the service
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				stmt, err := DB.Prepare("INSERT INTO Condition " +
					"(Username, MaxCalories, Diabetic, Halal, Vegan) " +
					"VALUES (?, ?, ?, ?, ?))")
				// Convert JSON to object
				json.Unmarshal(reqBody, &userC)
				// Open database and close it later
				db := openDB()
				defer db.Close()
				// Execute Query
				results, err := db.Exec(userC.Username, userC.MaxCalories, userC.Diabetic, userC.Halal, userC.Vegan)
				if err != nil {
					// Send error to client
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Username"))
					return
				}
				if rows, _ := results.RowsAffected(); rows > 0 {
					// Send success to client
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - " + strconv.FormatInt(rows, 10) + " row(s) affected"))
				}
			}
		}
	}
}
func (userC *UserCond) Update(w http.ResponseWriter, r *http.Request) {
	// Check valid key
	if !validKey(r) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Invalid key"))
		return
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "PUT" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				stmt, err := DB.Prepare("UPDATE Condition " +
					"SET MaxCalories=?, Diabetic=?, Halal=?, Vegan=? " +
					"WHERE Username=?")
				json.Unmarshal(reqBody, &userC)
				db := openDB()
				defer db.Close()
				results, err := db.Exec(userC.MaxCalories, userC.Diabetic, userC.Halal, userC.Vegan, userC.Username)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("404 - Not found"))
					return
				}
				if rows, _ := results.RowsAffected(); rows > 0 {
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - " + strconv.FormatInt(rows, 10) + " row(s) affected"))
				} else {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("404 - Not found"))
				}
			}
		}
	}
}
func (userC UserCond) Get(w http.ResponseWriter, r *http.Request) {
	// Check valid key
	if !validKey(r) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Invalid key"))
		return
	}
	// Get variables from client request
	params := mux.Vars(r)
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "GET" {
			query := "SELECT * " +
				"FROM MemberType " +
				"WHERE Username=?"
			db := openDB()
			defer db.Close()
			results, err := db.Query(query, params["Username"])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - No course found"))
				return
			}
			if results.Next() {
				err = results.Scan(&userC.Username, &userC.MaxCalories, &userC.Diabetic, &userC.Halal, &userC.Vegan)
				if err != nil {
					panic(err.Error())
				}
				//json.NewEncoder(w).Encode("ID: " + course.ID + ", Title: " + course.Title)
			}
		}
	}
}
*/
