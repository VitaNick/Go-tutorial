package main

import (
	"fmt"
	"organization"
)

func main() {
	p := organization.NewPerson("Nick", "Vanden Eynde", organization.NewEuropeanUnionIdentifier("123-45-6789", "Belgium"))
	err := p.SetTwitterHandler(organization.TwitterHandler("@ninjawulf98"))
	// fmt.Printf("%T\n", organization.TwitterHandler("test"))

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}

	// p.First = "Collin"
	// println(p.First)
	// println(p.Name.First)
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectUrl())
	// println(p.ID())
	// println(p.Country())
	// println(p.FullName())

	name1 := Name{First: "", Last: ""}
	// name2 := OtherName{First: "James", Last: "Wilson"}

	// ssn := organization.NewSocialSecurityNumber("123-45-6789")
	// eu := organization.NewEuropeanUnionIdentifier("123-45-6789", "Belgium")
	// eu2 := organization.NewEuropeanUnionIdentifier("123-45-6789", "Belgium")
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}

	// if name1 == (Name{}) {
	// 	println("We match")
	// }

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eu)

}

type Name struct {
	First  string
	Last   string
	Middle []string
}

type OtherName struct {
	First string
	Last  string
}
