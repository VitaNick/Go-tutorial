package main

import (
	"fmt"
	"reflect"
)

type Person string

const (
	Customer = "customer"
	Employee = "employee"
)

func main() {
	type employee struct {
		personId  int64
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmployee := employee{0, "Fred", "Flintstone"}

	newCustomer := customer{234, "Barney", "Rubble", "Slate Rock and Gravel"}

	addPerson(newEmployee)
	addPerson(newCustomer)

}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case Customer:
			empSqlString := "INSERT INTO employees (personId, firstName, lastName) VALUES (?,?,?)"

			fmt.Printf("SQL: %s\n", empSqlString)
			fmt.Printf("Added %v\n", v.Field(1))

		case Employee:
			empSqlString := "INSERT INTO customers (customerId, firstName, lastName, company) VALUES (?,?,?,?)"

			fmt.Printf("SQL: %s\n", empSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		}

		fmt.Println("It was a struct!")

		return true
	}

	return false
}
