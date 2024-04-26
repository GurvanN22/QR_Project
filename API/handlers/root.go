package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// This endpoint is a bit particular because it will be the root of our API and the /image endpoint because with golang servers when a url is
// not found it will redirect to the root of the server. So we will have to check if the url is /image and if it is we will redirect to the
//Image handler. If it is not we will return a 404 error.
func Root(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	// We check if the url start with /image
	if len(url) >= 6 {
		if url[:6] == "/image" {
			Image(w, r)
			return
		} else {
			Error404(w, r)
			return

		}
	} else {
		Error404(w, r)
		return
	}
}

func Error404(w http.ResponseWriter, r *http.Request) {
	response := ErrorResponse{
		Code:    404,
		Message: "Not Found for mor information go to /api",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
