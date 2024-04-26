package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type ImageInfoResponse struct {
	Image_id   string `json:"image_id"`
	User_id    string `json:"user_id"`
	Link       string `json:"link"`
	Created_at string `json:"created_at"`
}

// Info_image handles retrieval of image information by user ID.
// @Summary Get image information
// @Description Get information about images by user ID
// @Tags images
// @Accept  json
// @Produce  json
// @Param id query string true "User ID to retrieve image information"
// @Success 200 {array} ImageInfoResponse "Image information"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "No data found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /info-image [get]
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

	var result []ImageInfoResponse
	for rows.Next() {

		var image_id, user_id, link, created_at string
		var res ImageInfoResponse
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
		response := ErrorResponse{
			Message: "no data found",
			Code:    404,
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
