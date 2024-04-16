package data_functions

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CheckDataPath(path string, fill string) {
	// this function check if the data base is there and if not create it
	//!work in progress
	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("🔧 Database not found, creating it...")
		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		db, err := sql.Open("sqlite3", path)
		if err != nil {
			print(err)
		}
		defer db.Close()

		file, _ := os.ReadFile("db/query/creation.sql") //We open the help file in DB
		filef := string(file)
		_, err = db.Exec(filef)
		if err != nil {
			fmt.Println(err)
		}

		if fill == "true" {
			fill_file, _ := os.ReadFile("db/query/exemple.sql") //We open the help file in DB
			fileString := string(fill_file)
			_, err = db.Exec(fileString)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("✅ Database created")
	}
	if err == nil {
		fmt.Println("✅ Database connect successfully")
	}
}
