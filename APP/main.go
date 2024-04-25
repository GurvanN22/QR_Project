package main

import (
	"app/src"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	assets := http.FileServer(http.Dir("./Template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	fmt.Println("http://localhost" + port + " 🚀")
	// Page
	http.HandleFunc("/loginCookie", src.LoginHandler)
	http.HandleFunc("/", src.Home)
	http.HandleFunc("/login", src.Login)
	http.HandleFunc("/register", src.Register)
	http.HandleFunc("/createqr", src.CreateQrCode)
	http.HandleFunc("/registerqr", src.RegisterQR)
	http.HandleFunc("/listeQR", src.ListeQR)
	http.HandleFunc("/Profile", src.Profile)
	// API call
	http.HandleFunc("/createUser", src.Create_user)
	http.ListenAndServe(port, nil)
}
