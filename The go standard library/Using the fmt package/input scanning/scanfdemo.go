package main

import "fmt"

func main() {
	var firstName, lastName string
	fmt.Println("What is your name?")
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hello %s %s! Nice to meet you!\n", firstName, lastName)
}
