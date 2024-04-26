package main

import (
	"app/src"
	"fmt"
	"net/http"
	"os"
)

const port = ":8080"

func main() {
	staticDir := "./static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		os.Mkdir(staticDir, os.ModePerm)
	}

	assets := http.FileServer(http.Dir("./Template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	fmt.Println("http://localhost" + port + " 🚀")
	// Page
	http.HandleFunc("/", src.Home)
	http.HandleFunc("/login", src.Login)
	http.HandleFunc("/register", src.Register)
	http.HandleFunc("/createqr", src.CreateQrCode)
	http.HandleFunc("/registerqr", src.RegisterQR)
	http.HandleFunc("/listeQR", src.ListeQR)
	http.HandleFunc("/Profile", src.Profile)
	// call
	http.HandleFunc("/loginCookie", src.LoginHandler)
	http.HandleFunc("/createUser", src.Create_user)
	http.HandleFunc("/submit", src.SubmitLinkQR)
	http.ListenAndServe(port, nil)
}
