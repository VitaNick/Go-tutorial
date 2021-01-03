package main

import "organization"

func main() {
	p := organization.Person{FirstName: "James", LastName: "Wilson"}
	p.FirstName = "Collin"
	println(p.ID())
}
