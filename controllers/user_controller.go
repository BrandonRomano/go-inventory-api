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

	// Loading user from DB
	user := new(models.User)
	user.Username = username
	err := user.Load()

	// Failed to load user
	if err != nil {
		response.Success = false
		return
	}

	// Preparing authentication
	auth := new(models.Auth)
	auth.PrepareAuth(*user)

	// Success + Pass content
	response.Success = true
	response.Content = auth
}

func (c *UserController) AuthenticateUser(writer http.ResponseWriter, request *http.Request) {
	// Creating the response
	response := new(models.Response)
	defer response.PrintJSON(writer)

	// Grabbing form values
	username := request.FormValue("username")
	inputHash := request.FormValue("hash")

	// Loading User
	user := new(models.User)
	user.Username = username
	user.Load()

	// Checking if valid hash
	if inputHash != user.Hash {
		response.Success = false
		return
	}

	// Generating a session
	session := new(models.Session)
	session.UserId = user.Id
	session.Token = models.GenerateToken()
	session.Expiration = models.GetExpiration()
	session.Store()

	// Successful
	response.Content = session
	response.Success = true
}
