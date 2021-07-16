package main

import (
	"github.com/alexeipancratov/inventory-service/receipt"
	"log"
	"net/http"

	"github.com/alexeipancratov/inventory-service/database"
	"github.com/alexeipancratov/inventory-service/product"
	_ "github.com/go-sql-driver/mysql" // we're just importing for side effects, we're not using it directly
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	// starting the server mux (the default one)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
