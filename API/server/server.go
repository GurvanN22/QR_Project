package server

import (
	"api/handlers"
	"fmt"
	"net/http"
)

// @title API of the QRCode project
// @version 1.0
// @description This api handle the database of the QRCode project , it mean the users and the images stored in the database
// @host localhost:4000
// @BasePath /
func Start_api(port string) {

	// We create our endpoints
	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/api", handlers.Documentation)
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
