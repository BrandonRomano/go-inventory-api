package main

import (
	"github.com/brandonromano/inventory_api/database"
	"net/http"
	"os"
)

func main() {
	// Open + Closing DB
	database.OpenDatabase()
	defer database.CloseDatabase()

	// Setting up router + server
	prepareRouter()
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
