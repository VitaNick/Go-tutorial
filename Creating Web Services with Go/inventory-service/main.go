package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/webservices/database"
	"github.com/pluralsight/webservices/product"
)

const apiBasePath = "api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:5000", nil)
}
