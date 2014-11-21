package main

import (
	"net/http"
)

func main() {
	// Open + Closing DB
	openDatabase()
	defer closeDatabase()

	// Setting up router + server
	prepareRouter()
	http.ListenAndServe("localhost:4000", nil)
}
