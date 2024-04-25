package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func Info_image(w http.ResponseWriter, r *http.Request) {
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
	rows, err := db.Query("SELECT id, user_id , link , created_at FROM qrcode WHERE user_id = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var image_id, user_id, link, created_at string
	if rows.Next() {

		err := rows.Scan(&image_id, &user_id, &link, &created_at)
		if err != nil {
			log.Fatal(err)
		}
		// Create a response JSON object
		response := struct {
			Image_id   string `json:"image_id"`
			User_id    string `json:"user_id"`
			Link       string `json:"link"`
			Created_at string `json:"created_at"`
		}{
			Image_id:   image_id,
			User_id:    user_id,
			Link:       link,
			Created_at: created_at,
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
