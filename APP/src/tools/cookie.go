package tools

import (
	"net/http"
	"time"
)

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
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
