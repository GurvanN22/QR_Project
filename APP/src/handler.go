package src

import (
	"app/src/tools"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	qrcode "github.com/skip2/go-qrcode"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", nil)
}

func CreateQrCode(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	renderTemplate(w, "createQrcode.html", nil)
}
func RegisterQR(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	renderTemplate(w, "registerQR.html", nil)
}
func ListeQR(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	renderTemplate(w, "listeQR.html", nil)
}
func Profile(w http.ResponseWriter, r *http.Request) {
	tools.CheckCookie(w, r)
	renderTemplate(w, "profile.html", nil)
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

	http.Redirect(w, r, "/listeQR", http.StatusSeeOther)
}

// RenderTemplate & TemplateCache
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	var myCache, err = createTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	t, ok := myCache[tmpl]
	if !ok {
		fmt.Println("Could not get template from cache")
	}

	buffer := new(bytes.Buffer)
	t.Execute(buffer, p)
	buffer.WriteTo(w)
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("Template/page/*.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("Template/layout/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseFiles(matches...)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
