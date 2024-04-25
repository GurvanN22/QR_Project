package src

import (
	"app/src/tools"
	"fmt"
	"net/http"
	"time"
)

func Create_user(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("http://localhost:4000/createUser")
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/login?message=user+created", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	token := tools.GenerateToken()

	expiration := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Path:    "/",
		Expires: expiration,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
