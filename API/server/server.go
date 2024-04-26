package server

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func Start_api(port string) {

	// We create the server images endpoint
	styles := http.FileServer(http.Dir("db/images/"))
	http.Handle("/static/img/", http.StripPrefix("/static/img", styles))

	// We create our endpoints
	//http.HandleFunc("/", nil)
	//check the request method

	// We create the endpoint to create a user
	http.HandleFunc("/create-user", handlers.Create_user)
	http.HandleFunc("/connect-user", handlers.Connect_user)
	http.HandleFunc("/info-user", handlers.Info_user)

	http.HandleFunc("/new-image", handlers.New_image)
	http.HandleFunc("/info-image", handlers.Info_image)
	http.HandleFunc("/delete-image", handlers.Delete_image)

	// We start the server
	fmt.Println("✅ Server running on port :", port)
	http.ListenAndServe(":"+port, nil)
}
