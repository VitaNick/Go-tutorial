package main

import (
	"fmt"
	"organization"
)

func main() {
	p := organization.NewPerson("Nick", "Vanden Eynde", organization.NewSocialSecurityNumber("123-45-6789"))
	err := p.SetTwitterHandler(organization.TwitterHandler("@ninjawulf98"))
	fmt.Printf("%T\n", organization.TwitterHandler("test"))

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}

	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.FullName())
}
