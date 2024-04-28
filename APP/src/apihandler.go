package src

import (
	"app/src/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type CreateUserRequest struct {
	Pseudo   string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Create_user(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}
	name := r.Form.Get("pseudo")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// Créer une struct pour les données de l'utilisateur
	userData := CreateUserRequest{
		Pseudo:   name,
		Email:    email,
		Password: password,
	}

	// Convertir les données en JSON
	jsonData, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	// Envoyer une requête POST à votre API pour créer un utilisateur
	resp, err := http.Post("http://localhost:4000/create-user", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Erreur lors de l'envoi de la requête à l'API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Erreur lors de la création de l'utilisateur sur l'API", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers une page de confirmation
	http.Redirect(w, r, "/login?message=user+created", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		_, err := tools.AuthenticateUser(email, password)
		if err != nil {
			fmt.Println("Erreur d'authentification 1:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// cookie
		tools.SetCookie(w, r)
		http.Redirect(w, r, "/cookieHandler", http.StatusSeeOther)
		return
	}
}

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tools.SendCookieToAPI(r)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UploadImageFromForm(r *http.Request, apiURL string) error {
	// Récupérer le fichier image du formulaire
	file, header, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du fichier du formulaire: %v", err)
	}
	defer file.Close()

	// Créer un buffer pour stocker l'image
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Créer une partie pour l'image dans le formulaire multipart
	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la partie du formulaire: %v", err)
	}

	// Copier le contenu de l'image dans la partie du formulaire
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("erreur lors de la copie du contenu du fichier: %v", err)
	}

	// Fermer le formulaire multipart
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("erreur lors de la fermeture du formulaire multipart: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("réponse de l'API non réussie: %s", resp.Status)
	}

	return nil
}
func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := UploadImageFromForm(r, "http://localhost:4000/new-image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'envoi de l'image à l'API: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "L'image a été envoyée avec succès à l'API.")
}
