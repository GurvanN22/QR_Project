package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type UserInfoResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	DATA    struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password int    `json:"password"`
	}
}

// Info_user handles retrieval of user information by ID.
// @Summary Get user information
// @Description Get information about a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query string true "User ID to retrieve information"
// @Success 200 {object} UserInfoResponse "User information"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /info-user [get]
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
		response := ErrorResponse{
			Message: "wrong field : id",
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
		response := UserInfoResponse{
			Message: "user found",
			Code:    200,
		}
		response.DATA.Name = name
		response.DATA.Email = email
		response.DATA.Password = len(password)

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
}
func GetidByEamilH(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	id, err := GetIDByEmail(email)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID de l'utilisateur", http.StatusInternalServerError)
		return
	}
	if id == "" {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}
	w.Write([]byte(id))
}
func GetIDByEmail(email string) (string, error) {
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT id FROM user WHERE email = ?;", email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // User not found
		}
		return "", err
	}

	return id, nil
}

func GetIdByCookie(cookie string) (string, error) {
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT user_id FROM session_cookie WHERE session_id = ?;", cookie).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // User not found
		}
		return "", err
	}

	return id, nil
}

func GetIdByCookieH(w http.ResponseWriter, r *http.Request) {
	cookie := r.FormValue("cookie")
	id, err := GetIdByCookie(cookie)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID de l'utilisateur", http.StatusInternalServerError)
		return
	}
	if id == "" {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}
	w.Write([]byte(id))
}
