package tools

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func GenerateToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	tokenString := base64.StdEncoding.EncodeToString(token)
	return tokenString
}

func CheckCookie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
