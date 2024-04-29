package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Code    int
	Message string
}

// Delete_image handles deletion of an image by ID.
// @Summary Delete an image
// @Description Delete an image by ID
// @Tags images
// @Accept  json
// @Produce  json
// @Param id query string true "Image ID to delete"
// @Success 200 {object} Response "Image deleted"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /delete-image [delete]
func Delete_image(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodDelete(w, r) {
		return
	}

	id := r.FormValue("id")

	// We get the data from the request
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q, err := db.Prepare("DELETE FROM qrcode WHERE ? = id;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.Exec(id)
	if err != nil {
		response := Response{
			Code:    404,
			Message: "Image not found",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		// Write the JSON response to the http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	// We delete the image from db/images folder
	err = os.Remove("db/images/" + id + ".png")
	if err != nil {
		log.Fatal(err)
	}

	response := Response{
		Code:    200,
		Message: "Image deleted",
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
