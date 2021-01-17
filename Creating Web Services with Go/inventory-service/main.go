package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/webservices/database"
	"github.com/pluralsight/webservices/product"
	"github.com/pluralsight/webservices/receipt"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	http.ListenAndServe("localhost:5000", nil)
}
