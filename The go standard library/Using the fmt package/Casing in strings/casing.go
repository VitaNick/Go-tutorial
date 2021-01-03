package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		sampleString := "Never trust a programmer who carries a screwdriver\n"

		fmt.Println("Before: " + sampleString)

		strLowerCase := strings.ToLower(sampleString)
		fmt.Println("Lower case: " + strLowerCase)

		strUpperCase := strings.ToUpper(sampleString)
		fmt.Println("Upper case: " + strUpperCase)

		strTitleCase := strings.Title(sampleString)
		fmt.Println("Title case: " + strTitleCase)
	}

	{
		sampleString := "welcome to the dollhouse\n"

		fmt.Println(strings.Title(sampleString))
		fmt.Println(properTitle(sampleString))
	}
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallWords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallWords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")
}
