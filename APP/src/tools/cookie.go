package tools

import (
	"fmt"
	"net/http"
	"time"
)

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("session_token")
	if err != nil {
		if err != http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		return true
	}
	return true
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	token := GenerateToken()
	expiration := time.Now().Add(24 * time.Hour)
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   token,
		Path:    "/",
		Expires: expiration,
	}
	http.SetCookie(w, cookie)
}

func SendCookieToAPI(r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Printf("no cookie found")
		}
		req, err := http.NewRequest("POST", "http://localhost:4000/cookies", nil)
		if err != nil {
			fmt.Printf("erreur lors de la création de la demande : %v", err)
			return
		}
		req.AddCookie(cookie)

	}
}
