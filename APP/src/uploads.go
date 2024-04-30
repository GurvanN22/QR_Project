package src

import (
	"app/src/tools"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}
	link := r.FormValue("link")
	userId, err := tools.GetIdByCookie(w, r)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID utilisateur", http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("link", link)
	_ = writer.WriteField("user_id", strconv.Itoa(userId))
	fileWriter, _ := writer.CreateFormFile("file", "image.png")
	_, _ = io.Copy(fileWriter, &buf)
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:4000/new-image", body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to upload image", resp.StatusCode)
		return
	}
	fmt.Print(w, "Image uploaded successfully!")
}

func SubmitLinkQR(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	link := r.FormValue("link")

	if link == "" {
		http.Error(w, "Le lien est vide", http.StatusBadRequest)
		return
	}

	qrCode, err := qrcode.Encode(link, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du QR code", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("static/qrcode.png", qrCode, 0644)
	if err != nil {
		http.Error(w, "Erreur lors de l'enregistrement du QR code", http.StatusInternalServerError)
		return
	}

	// Get form data
	userId, err := tools.GetIdByCookie(w, r)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID utilisateur", http.StatusInternalServerError)
		return
	}
	// Retrieve the uploaded file
	file, err := os.Open("static/qrcode.png")
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Buffer to store the file data
	var buf bytes.Buffer

	// Copy file data to buffer
	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Write form fields
	_ = writer.WriteField("link", link)
	_ = writer.WriteField("user_id", strconv.Itoa(userId))

	// Write file field
	fileWriter, _ := writer.CreateFormFile("file", "image.png")
	_, _ = io.Copy(fileWriter, &buf)

	// Close multipart writer
	writer.Close()

	// Send POST request to API
	apiURL := "http://localhost:4000/new-image" // Change this to your API endpoint
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to upload image", resp.StatusCode)
		return
	}

	// Display success message
	http.Redirect(w, r, "/list", http.StatusSeeOther)
	os.Remove("static/qrcode.png")

}
