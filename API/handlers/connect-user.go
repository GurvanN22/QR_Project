package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Create a response JSON object
type ConnectUserResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Connect_user handles user authentication.
// @Summary Authenticate user
// @Description Authenticate user with provided email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Success 200 {object} ConnectUserResponse "Successfully authenticated"
// @Failure 400 {object} ErrorResponse "Missing fields"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/connect [post]
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
		response := ErrorResponse{
			Message: "missing fields",
			Code:    400,
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

	// We encrypt the password
	password = tools.Chiffrement(email, password)

	// We get the name, email, and password from the database
	rows, err := db.Query("SELECT id FROM user WHERE email = ? AND password = ?;", email, password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	} else {

		var response = ErrorResponse{
			Message: "user not found",
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
	// Create a response JSON object

	type ConnectUserResponse struct {
		Message string `json:"message"`
		Id      string `json:"id"`
		Code    int    `json:"code"`
	}

	var response = ConnectUserResponse{
		Message: "user authenticated",
		Id:      id,
		Code:    200,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
