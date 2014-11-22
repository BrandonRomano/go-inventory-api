package controllers

import (
	"fmt"
	"github.com/brandonromano/inventory_api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (product *ProductController) Get(writer http.ResponseWriter, request *http.Request) {
	// Getting the ID
	vars := mux.Vars(request)
	id := vars["id"]

	// Loading the product model
	productModel := new(models.Product)
	productModel.Load(id)

	fmt.Fprintln(writer, productModel.ToJSON())
}

func (product *ProductController) Post(writer http.ResponseWriter, request *http.Request) {
	// Formatting price string
	priceString := request.FormValue("price")
	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		panic(err) // Invalid input
	}

	// Creating products model
	productModel := new(models.Product)
	productModel.Name = request.FormValue("name")
	productModel.Description = request.FormValue("description")
	productModel.Price = float32(price)

	// Storing products model
	productModel.Store()

	// TODO print response
}
