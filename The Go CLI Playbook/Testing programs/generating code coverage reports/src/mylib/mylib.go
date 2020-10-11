package mylib

import "fmt"

func adder(l, r int) int {
	fmt.Printf("Adding %v and %v", l, r)
	return l + r
}

func subtractor(l, r int) int {
	return l - r
}
