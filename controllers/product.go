package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct{}

func (product *Product) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		product.get(writer, mux.Vars(request))
	}
}

func (product *Product) get(writer http.ResponseWriter, vars map[string]string) {
	fmt.Fprintln(writer, vars["id"])
}
