package main

import "fmt"

func main() {
	{
		ourString := "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6f\x65\x21"

		for i := 0; i < len(ourString); i++ {
			fmt.Printf("%q\n ", ourString[i])
		}

	}

	{
		ourString := "\xbd\xb2\x3f\xbc\x20\xe2\x8c\x98"

		for i := 0; i < len(ourString); i++ {
			fmt.Printf("%q\n ", ourString[i])
		}
	}

	newString := "This is a string"
	fmt.Println(newString[3])
	fmt.Println(newString[0:5])

}
