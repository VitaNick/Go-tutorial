package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	personId  int
	firstname string
	lastname  string
}

func main() {

	employees := make([]employee, 3)

	employees = append(employees, employee{0, "Lloyd", "Christmas"})
	employees = append(employees, employee{1, "Harry", "Dunne"})
	employees = append(employees, employee{2, "Sea", "Bass"})

	// name
	fmt.Printf("The name of type is %v\n", reflect.TypeOf(employees).Name())
	// type
	fmt.Printf("The type is %v\n", reflect.TypeOf(employees))
	// kind
	fmt.Printf("The kind is %v\n", reflect.TypeOf(employees).Kind())
	// value
	fmt.Printf("The value is %v\n", reflect.ValueOf(employees))

	eType := reflect.TypeOf(employees)

	newEmployeeList := reflect.MakeSlice(eType, 0, 0)

	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employee{3, "Marry", "Swanson"}))

	fmt.Printf("First list of employees: %v\n\n", employees)

	fmt.Printf("List created by reflection: %v\n", newEmployeeList)

}
