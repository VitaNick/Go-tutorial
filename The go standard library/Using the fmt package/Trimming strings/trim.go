package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		sampleString := "     This is our text     "

		fmt.Printf("%q\n", sampleString)

		newString := strings.TrimSpace(sampleString)

		fmt.Printf("%q\n", newString)
	}

	{
		sampleString := "https://www.pluralsight.com"

		domainName := strings.TrimPrefix(sampleString, "https://")

		fmt.Println(domainName)
	}

	{
		sampleString := "test.doc"

		newFileName := strings.TrimSuffix(sampleString, ".doc")

		fmt.Println(newFileName)
	}
}
