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
	productController := new(controllers.ProductController)
	router.HandleFunc("/products/{id:[0-9]+}", productController.Get).Methods("GET")
	router.HandleFunc("/products", productController.Post).Methods("POST")

	// User Controller
	userController := new(controllers.UserController)
	router.HandleFunc("/users/auth/{username}", userController.GetAuthentication).Methods("GET")

	// Handle the router
	http.Handle("/", router)
}
