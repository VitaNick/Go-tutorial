package main

import (
	"fmt"
	"strings"
)

func main() {
	ourString := "This is a string!"

	stringCollection := strings.SplitN(ourString, " ", 2)

	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}
}
