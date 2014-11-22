package main

import (
	"github.com/brandonromano/inventory_api/database"
	"net/http"
)

func main() {
	// Open + Closing DB
	database.OpenDatabase()
	defer database.CloseDatabase()

	// Setting up router + server
	prepareRouter()
	http.ListenAndServe("localhost:4000", nil)
}
