package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func Create_user(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}
	// We get the data from the request
	name := r.FormValue("pseudo")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// We check if the fields are not empty
	if name == "" || email == "" || password == "" {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "missing field : pseudo, email or password",
			CODE:    400,
		}

		// Convert the response object to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}

		// Write the JSON response to the http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO user (name,email,password) VALUES (?,?,?);", name, email, password)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	// Create a response JSON object
	response := struct {
		Message string `json:"message"`
		ID      int64  `json:"id"`
		CODE    int    `json:"code"`
	}{
		Message: "user created",
		ID:      id,
		CODE:    200,
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}