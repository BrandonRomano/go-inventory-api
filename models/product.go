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
