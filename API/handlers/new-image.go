package handlers

import (
	"api/handlers/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func New_image(w http.ResponseWriter, r *http.Request) {
	// We check the request method
	if !tools.CheckRequestMethodPost(w, r) {
		return
	}
	// We get the img from the request
	link := "link.com"
	user_id := 1
	// Parse the multipart form to get the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}

	// Retrieve the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	///////////////////////////////////////////////
	// Generate a random filename

	id := random_word(16)

	// Create a new file in the images folder
	newFile, err := os.Create("db/images/" + id + ".png")
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the contents of the uploaded file to the new file

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	///////////////////////////////////////////////
	// Open SQLite database file
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	created_at := time.Now()

	// Insert the image data into the database
	_, err = db.Exec("INSERT INTO qrcode (id , user_id , link , created_at) VALUES (?, ?,?,?)", id, user_id, link, created_at)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to store image in database", http.StatusInternalServerError)
		return
	}

	///////////////////////////////////////////////
	// Return success
	type Response struct {
		Message string
		Code    int
		ID      string
	}

	response := Response{
		Message: "Image added successfully",
		Code:    200,
		ID:      id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON response to the http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func random_word(length int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bytes)
}
