package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct{}

func (product *Product) Get(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(writer, vars["id"])
}
