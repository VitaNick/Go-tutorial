package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go Language. It's one of my favorites"

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		result := strings.HasPrefix(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string starts with %s!\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT start with %s\n", searchTerm)
		}
	}

}
