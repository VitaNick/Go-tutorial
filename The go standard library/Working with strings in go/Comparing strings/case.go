package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "This is a string!"
	string2 := "this is a string!"

	fmt.Println(CompareCaseIns(string1, string2))
}

func CompareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}
	}

	return false
}
