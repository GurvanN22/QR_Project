package handlers

import (
	"api/handlers/tools"
	"net/http"
)

func Create_user(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}

	// We get the data from the request
	//email := r.FormValue("email")
	//password := r.FormValue("password")

}
