package main

import (
	"fmt"
	"net/http"

	"github.com/ninjawulf98/webservice/controllers"
)

func main() {
	fmt.Println("Starting http server...")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
