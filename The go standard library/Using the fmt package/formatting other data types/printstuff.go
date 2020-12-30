package main

import "fmt"

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := point{2, 3}
	fmt.Printf("%v\n", p)

	newPerson := Person{"Jeremy", "Morgan", 42}
	fmt.Printf("%v\n", newPerson)

	fmt.Printf("%b\n", 4567)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%c\n", 42)
	fmt.Printf("%x\n", 455)

}
