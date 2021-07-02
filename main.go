package main

import (
	"net/http"

	"github.com/alexeipancratov/inventory-service/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	// starting the server mux (the default one)
	http.ListenAndServe(":5000", nil)
}
