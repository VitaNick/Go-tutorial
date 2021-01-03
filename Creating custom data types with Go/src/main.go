package main

import "organization"

func main() {
	p := organization.NewPerson("Collin", "Wilson")

	println(p.ID())
	println(p.FullName())
}
