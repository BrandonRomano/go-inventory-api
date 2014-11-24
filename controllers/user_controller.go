package controllers

import (
	"github.com/brandonromano/inventory_api/models"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct{}

func (c *UserController) GetAuthentication(writer http.ResponseWriter, request *http.Request) {
	// Creating the response
	response := new(models.Response)
	defer response.PrintJSON(writer)

	// Getting the Username
	vars := mux.Vars(request)
	username := vars["username"]

	// Requesting authentication
	auth := new(models.Auth)
	err := auth.RequestAuthentication(username)

	// Request failed
	if err != nil {
		response.Success = false
		return
	}

	// Success + Pass content
	response.Success = true
	response.Content = auth
}
