package server

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func Start_api(port string) {
	// We create our endpoints
	http.HandleFunc("/", nil)
	//check the request method

	// We create the endpoint to create a user
	http.HandleFunc("/createUser", handlers.Create_user)

	// We start the server
	http.ListenAndServe(":"+port, nil)
	fmt.Println("✅ Server running on port", port)
}
