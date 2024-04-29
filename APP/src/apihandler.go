package src

import (
	"app/src/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Send the POST request to the API with the form values as query parameters
	resp, err := http.Post(fmt.Sprintf("http://localhost:4000/create-user?name=%s&email=%s&password=%s", name, email, password), "", nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'envoi de la requête à l'API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse de l'API", http.StatusInternalServerError)
		return
	}

	// Create a struct to store the response data
	var response struct {
		Message string `json:"message"`
	}

	// Unmarshal the response JSON into the struct
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion de la réponse en JSON", http.StatusInternalServerError)
		return
	}

	// Print the message from the response
	fmt.Println(response.Message)

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
		idResp, err := http.Get(fmt.Sprintf("http://localhost:4000/get-id?email=%s", email))
		if err != nil {
			http.Error(w, "Erreur lors de l'envoi de la requête à l'API", http.StatusInternalServerError)
			return
		}
		defer idResp.Body.Close()
		idData, err := ioutil.ReadAll(idResp.Body)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture de la réponse de l'API", http.StatusInternalServerError)
			return
		}
		id, err := strconv.Atoi(string(idData))
		if err != nil {
			http.Error(w, "Erreur lors de la conversion de l'ID en entier"+string(idData), http.StatusInternalServerError)
			return
		}
		url := "http://localhost:8080/cookieHandler?id=" + strconv.Itoa(id)
		// cookie
		tools.SetCookie(w, r)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}
}

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	cookie, err := r.Cookie("session_id")
	fmt.Printf("http://localhost:4000/cookie-save?session=%s&user_id=%s", cookie.Value, id)
	if err != nil {
		fmt.Fprintf(w, "Cookie not found: %v", err)
		return
	}
	resp, err := http.Post(fmt.Sprintf("http://localhost:4000/cookie-save?session=%s&user_id=%s", cookie.Value, id), "", nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'envoi de la requête à l'API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse de l'API", http.StatusInternalServerError)
		return
	}
	var response struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion de la réponse en JSON", http.StatusInternalServerError)
		return
	}
	fmt.Println(response.Message)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Erreur lors de la création de l'utilisateur sur l'API", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
