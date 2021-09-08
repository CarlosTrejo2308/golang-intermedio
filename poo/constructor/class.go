package main

import "fmt"

type Employee struct {
	age      int
	name     string
	vacation bool
}

func NewEmployee(name string, age int, vacation bool) *Employee {
	emp := Employee{age, name, vacation}
	return &emp
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%+v\n", e)

	// 2
	e = Employee{age: 42, name: "John", vacation: true}
	fmt.Printf("%+v\n", e)

	// 3
	e3 := new(Employee)
	fmt.Printf("%+v\n", *e3)
	e3.age = 42
	e3.name = "John"
	e3.vacation = true
	fmt.Printf("%+v\n", *e3)

	// 4
	e4 := NewEmployee("John", 42, true)
	fmt.Printf("%+v\n", *e4)
}
