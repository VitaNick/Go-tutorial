package main

import "fmt"

func main() {
	var age = 42
	var name = "Nick"
	_output := fmt.Sprint("My name is ", name, " and I am ", age, " years old\n")
	print(_output)
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	var pi float32 = 3.141592
	fmt.Printf("Pi is %.2f\n", pi)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 98.999, 12.3456, 12.01)

	output := fmt.Sprintf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	print(output)
	fmt.Printf("|%-7s|%-7s|%-7q|\n", "a", "ab", 100)

}
