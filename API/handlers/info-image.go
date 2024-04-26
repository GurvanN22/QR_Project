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

	type Response struct {
		Image_id   string
		User_id    string
		Link       string
		Created_at string
	}

	var result []Response
	for rows.Next() {

		var image_id, user_id, link, created_at string
		var res Response
		// We return all the images data link to the user by the id of the user , here we use a loop to get all the data
		err := rows.Scan(&image_id, &user_id, &link, &created_at)

		res.Created_at = created_at
		res.Image_id = image_id
		res.User_id = user_id
		res.Link = link

		result = append(result, res)

		if err != nil {
			log.Fatal(err)
		}
	}

	// We check if the user exists
	if len(result) == 0 {
		// Create a response JSON object
		response := struct {
			Message string `json:"message"`
			CODE    int    `json:"code"`
		}{
			Message: "no data found",
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

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
