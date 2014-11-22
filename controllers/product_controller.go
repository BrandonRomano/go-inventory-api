package controllers

import (
	"github.com/brandonromano/inventory_api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (product *ProductController) Get(writer http.ResponseWriter, request *http.Request) {
	response := new(models.Response)
	defer response.PrintJSON(writer)

	// Getting the ID
	vars := mux.Vars(request)
	id := vars["id"]

	// Loading the product model
	productModel := new(models.Product)
	err := productModel.Load(id)

	// Couldn't load product model
	if err != nil {
		// Couldn't load product model
		response.Success = false
		return
	}

	// Successful response
	response.Success = true
	response.Content = productModel
}

func (product *ProductController) Post(writer http.ResponseWriter, request *http.Request) {
	response := new(models.Response)
	defer response.PrintJSON(writer)

	// Formatting price string
	priceString := request.FormValue("price")
	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		response.Success = false
		response.Message = "invalid input"
		return
	}

	// Creating products model
	productModel := new(models.Product)
	productModel.Name = request.FormValue("name")
	productModel.Description = request.FormValue("description")
	productModel.Price = float32(price)

	// Storing products model
	id, store_error := productModel.Store()
	if store_error != nil {
		response.Success = false
		return
	}

	// Successful response
	response.Success = true
	productModel.Id = id
	response.Content = productModel
}
