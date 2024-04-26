package handlers

import (
	"api/handlers/tools"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Image(w http.ResponseWriter, r *http.Request) {

	if !tools.CheckRequestMethodGet(w, r) {
		return
	}

	// We get the entire url of the request
	id := r.URL.Path[7:]

	// We check if the file exist
	if _, err := os.Stat("db/images/" + id + ".png"); os.IsNotExist(err) {
		response := ErrorResponse{
			Code:    404,
			Message: "The image does not exist",
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

	// We serve the file
	http.ServeFile(w, r, "db/images/"+id+".png")
}
