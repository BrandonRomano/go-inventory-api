package main

import (
	"github.com/brandonromano/inventory_api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * This function prepares the router and is responsible
 * for linking all controller functions to URLs
 */
func prepareRouter() {
	// Creating a router
	router := mux.NewRouter()

	// Product Controller
	productController := new(controllers.Product)
	router.Handle("/products/{id:[0-9]+}", productController).
		Name("products").
		Methods("GET")

	// Handle the router
	http.Handle("/", router)
}