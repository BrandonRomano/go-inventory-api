package models

import (
	"fmt"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float32
}

func (product *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	html_str := "<h1>Product: %v</h1>" +
		"<p>Description: %v</p>" +
		"<strong>Price: %v</strong>"

	fmt.Fprintf(w, html_str, product.Name, product.Description, product.Price)
}
