package main

import (
	"api/data_functions"
	"api/handlers/tools"
	"api/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var database_path string

func main() {
	// We load the .env file
	err := godotenv.Load()
	if err != nil {
		panic("🔥 Error loading .env file, is the file there? the file must be in the same folder as the main.go file")
	} else {
		println("✅ The .env file has been loaded successfully")
	}
	// We extract the data from the .env file
	port := os.Getenv("PORT")
	fill := os.Getenv("DATA_FILL")
	database_path = os.Getenv("DATA_BASE_FILE")

	// We connect to the database
	data_functions.CheckDataPath(database_path, fill)

	fmt.Println(tools.Chiffrement("admin@exemple.com", "admin"))
	server.Start_api(port)
}
