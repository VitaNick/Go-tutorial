package main

import (
	"net/http"

	"github.com/pluralsight/webservices/product"
)

const apiBasePath = "api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:5000", nil)
}
