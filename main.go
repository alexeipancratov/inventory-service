package main

import (
	"net/http"

	"github.com/alexeipancratov/inventory-service/database"
	"github.com/alexeipancratov/inventory-service/product"
	_ "github.com/go-sql-driver/mysql" // we're just importing for side effects, we're not using it directly
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	// starting the server mux (the default one)
	http.ListenAndServe(":5000", nil)
}
