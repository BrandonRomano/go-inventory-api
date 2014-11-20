// Visit http://localhost:4000/ to see the greeting.
package main

import (
	"github.com/brandonromano/inventory_api/models"
	"github.com/gorilla/mux"
	"net/http"
)

func prepareRouter() {
	product := &models.Product{
		Name:        "Product One",
		Description: "Wow this is a description",
		Price:       12.50,
	}

	r := mux.NewRouter()
	r.Handle("/", product)
	http.Handle("/", r)
}

func main() {
	prepareRouter()
	http.ListenAndServe("localhost:4000", nil)
}
