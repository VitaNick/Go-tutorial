package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var firstName string
	var lastName string

	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hello %s %s nice to meet you!\n", firstName, lastName)
}
