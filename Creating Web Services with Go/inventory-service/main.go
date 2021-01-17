package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pluralsight/webservices/product"
)

const apiBasePath = "api"

var productList []Product

func init() {
	productsJSON := `[
		  {
    "productId": 1,
    "manufacturer": "Johns-Jenkins",
    "sku": "p5z343vdS",
    "upc": "939581000000",
    "pricePerUnit": "497.45",
    "quantityOnHand": 9703,
    "productName": "sticky note"
  },
  {
    "productId": 2,
    "manufacturer": "Hessel, Schimmel and Feeney",
    "sku": "i7v300kmx",
    "upc": "740979000000",
    "pricePerUnit": "282.29",
    "quantityOnHand": 9217,
    "productName": "leg warmers"
  },
  {
    "productId": 3,
    "manufacturer": "Swaniawski, Bartoletti and Bruen",
    "sku": "q0L657ys7",
    "upc": "111730000000",
    "pricePerUnit": "436.26",
    "quantityOnHand": 5905,
    "productName": "lamp shade"
  }]`

	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := 1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}

	return highestID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}

	return nil, 0
}

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:5000", nil)
}
