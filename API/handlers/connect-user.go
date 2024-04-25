package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Connect_user(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}
	// We get the data from the request
	email := r.FormValue("email")
	password := r.FormValue("password")

	// We check if the fields are not empty
	if email == "" || password == "" {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "missing field : email or password",
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

	var id string
	// We get the name, email, and password from the database
	rows, err := db.Query("SELECT id FROM user WHERE email = ? AND password = ?;", email, password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		fmt.Println(id)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "wrong email or password",
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
	// Create a response JSON object
	response := struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		CODE    int    `json:"code"`
	}{
		Message: "user connecte",
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
