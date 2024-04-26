package src

import (
	"app/src/tools"
	"net/http"
	"os"

	qrcode "github.com/skip2/go-qrcode"
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
	tools.CheckCookie(w, r)
	tools.RenderTemplate(w, "listeQR.html", nil)
}
func Profile(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	tools.RenderTemplate(w, "profile.html", nil)
}

func SubmitLinkQR(w http.ResponseWriter, r *http.Request) {
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

}
