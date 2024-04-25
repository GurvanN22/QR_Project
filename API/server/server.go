package server

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func Start_api(port string) {
	// We create our endpoints
	//http.HandleFunc("/", nil)
	//check the request method

	// We create the endpoint to create a user
	http.HandleFunc("/create-user", handlers.Create_user)
	http.HandleFunc("/info-user", handlers.Info_user)
	http.HandleFunc("/connect-user", handlers.Connect_user)

	// We start the server
	fmt.Println("✅ Server running on port :", port)
	http.ListenAndServe(":"+port, nil)
}
