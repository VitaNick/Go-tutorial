package main

import "fmt"

func main() {
	a := func(name string) string {
		fmt.Printf("My first %s function\n", name)

		return name
	}

	value := a("anonymous")
	println(value)
}
