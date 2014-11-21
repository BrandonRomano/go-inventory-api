package controllers

import (
	"fmt"
	"github.com/brandonromano/inventory_api/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct{}

func (product *Product) Get(writer http.ResponseWriter, request *http.Request) {
	// Getting the ID
	vars := mux.Vars(request)
	id := vars["id"]

	// Loading the product model
	productModel := new(models.Product)
	productModel.Load(id)

	fmt.Fprintln(writer, productModel.ToJSON())
}
