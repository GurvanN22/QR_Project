package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func Info_user(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodGet(w, r) {
		return
	}
	// We get the data from the request
	id := r.FormValue("id")

	// We check if the fields are not empty
	if id == "" {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "wrong field : id",
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

	// We get the name, email, and password from the database
	rows, err := db.Query("SELECT name, email, password FROM user WHERE id = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name, email, password string
	if rows.Next() {
		err := rows.Scan(&name, &email, &password)
		if err != nil {
			log.Fatal(err)
		}
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
			DATA    struct {
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password int    `json:"password"`
			}
		}{
			Message: "user's information",
			CODE:    200,
			DATA: struct {
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password int    `json:"password"`
			}{
				Name:     name,
				Email:    email,
				Password: len(password),
			},
		}
		// Convert the response object to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		// Write the JSON response to the http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "user not found",
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
}
