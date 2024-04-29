package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CreateUserResponse struct {
	Message string `json:"message"`
	ID      int64  `json:"id"`
	Code    int    `json:"code"`
}

// @Summary Create a new user
// @Description Create a new user with provided name, email, and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param pseudo formData string true "User's pseudo"
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Success 200 {object} CreateUserResponse "Successfully created user"
// @Failure 400 {object} ErrorResponse "Missing fields"
// @Failure 500 {object} ErrorResponse "Internal server error"s
// @Router /create-user [post]
func Create_user(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}
	// We get the data from the request
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// We check if the fields are not empty
	if name == "" || email == "" || password == "" {
		fmt.Println("missing fields")
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

	// We encrypt the password with a homemade encryption method
	password = tools.Chiffrement(email, password)

	db, err := sql.Open("sqlite3", "./db/data.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO user (name,email,password) VALUES (?,?,?);", name, email, password)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("error 2")
	}

	// Create a response JSON object
	response := CreateUserResponse{
		Message: "user created",
		ID:      id,
		Code:    200,
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
