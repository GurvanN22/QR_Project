package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SaveCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}
	r.ParseForm()

	session_id := r.FormValue("session")
	user_id := r.FormValue("user_id")
	if session_id == "" {
		http.Error(w, "Session ID manquant", http.StatusBadRequest)
		return
	}

	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Enregistrer le cookie de session dans la base de données
	_, err = db.Exec("INSERT INTO session_cookie (session_id , user_id) VALUES (?,?)", session_id, user_id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de l'enregistrement du cookie dans la base de données", http.StatusInternalServerError)
		return
	}
	// Create a response JSON object
	response := CreateUserResponse{
		Message: "cookie saved",
		Code:    200,
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
