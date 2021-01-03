package main

import (
	"fmt"
	"organization"
)

func main() {
	p := organization.NewPerson("Collin", "Wilson")
	err := p.SetTwitterHandler("@jam_wils")

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}

	println(p.ID())
	println(p.FullName())
}
