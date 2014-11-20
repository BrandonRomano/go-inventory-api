package main

import (
	"net/http"
)

func main() {
	prepareRouter()
	http.ListenAndServe("localhost:4000", nil)
}
