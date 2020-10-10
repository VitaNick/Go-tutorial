package main

import (
	"fmt"

	"github.com/ninjawulf98/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Nick",
		LastName:  "Lol",
	}
	fmt.Println(u)
}
