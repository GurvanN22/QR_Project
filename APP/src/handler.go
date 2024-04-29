package src

import (
	"app/src/tools"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tools.RenderTemplate(w, "home.html", nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	tools.RenderTemplate(w, "login.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	tools.RenderTemplate(w, "register.html", nil)
}

func CreateQrCode(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	tools.RenderTemplate(w, "createQrcode.html", nil)
}
func RegisterQR(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	tools.RenderTemplate(w, "registerQR.html", nil)
}
func ListeQR(w http.ResponseWriter, r *http.Request) {
	type ImageInfoResponse struct {
		ImageID   string `json:"image_id"`
		UserID    string `json:"user_id"`
		Link      string `json:"link"`
		CreatedAt string `json:"created_at"`
	}

	tools.CheckCookie(w, r)
	id, err := tools.GetIdByCookie(w, r)
	if err != nil {
		http.Error(w, "Échec de récupération de l'ID utilisateur", http.StatusInternalServerError)
		return
	}
	url := "http://localhost:4000/info-image?id=" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Échec de récupération des informations sur les images", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var images []ImageInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&images)
	if err != nil {
		http.Error(w, "Échec de décodage de la réponse JSON", http.StatusInternalServerError)
		return
	}
	for _, image := range images {
		imageURL := fmt.Sprintf("http://localhost:4000/image/%s", image.ImageID)
		resp, err := http.Get(imageURL)
		if err != nil {
			http.Error(w, "Échec de récupération de l'URL de l'image", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

	}
	data := struct {
		Images []ImageInfoResponse
	}{
		Images: images,
	}

	tools.RenderTemplate(w, "listeQR.html", data)

}
func Profile(w http.ResponseWriter, r *http.Request) {
	userId, err := tools.GetIdByCookie(w, r)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID utilisateur", http.StatusInternalServerError)
		return
	}

	url := "http://localhost:4000/info-user?id=" + strconv.Itoa(userId)
	resp, err := http.Get(url)
	type UserInfoResponse struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password int    `json:"password"`
		} `json:"DATA"`
	}

	if err != nil {
		http.Error(w, "Erreur lors de la récupération des informations utilisateur", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo UserInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse de l'API", http.StatusInternalServerError)
		return
	}

	data := struct {
		Name     string
		Email    string
		Password int
	}{
		Name:     userInfo.Data.Name,
		Email:    userInfo.Data.Email,
		Password: userInfo.Data.Password,
	}

	tools.RenderTemplate(w, "profile.html", data)

}
