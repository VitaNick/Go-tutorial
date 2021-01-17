package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const productsBasePath = "products"

func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(productsHandler)
	handleProduct := http.HandlerFunc(productHandler)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), handleProducts)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), handleProduct)

}

// Product
type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList := getProductList()
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(productsJson)

	case http.MethodPost:
		// add a new product to the list
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader((http.StatusBadRequest))
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := addOrUpdateProduct(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "/products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	product := getProduct(productID)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return a single product
		productJSON, err := json.Marshal(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)

	case http.MethodPut:
		// Update product in the list
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if updatedProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		addOrUpdateProduct(updatedProduct)
		w.WriteHeader(http.StatusOK)

		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}
