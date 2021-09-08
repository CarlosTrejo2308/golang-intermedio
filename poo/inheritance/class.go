package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Inheritance tipo anonimo
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemproaryEmployee struct {
	Person
	Employee
	taxRate float64
}

type PrintInfo interface {
	getMessage() string
}

// Composicion sobre herencia
func (ft FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a full time employee", ft.name, ft.age)
}

func (te TemproaryEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a temprary employee", te.name, te.age)
}

func GetMessage(pi PrintInfo) {
	fmt.Println(pi.getMessage())
}

func main() {
	f := FullTimeEmployee{}
	f.name = "John"
	f.age = 30
	f.id = 1
	f.endDate = "2019-01-01"
	fmt.Printf("%+v\n", f)

	t := TemproaryEmployee{}
	t.name = "Mary"
	t.age = 25
	t.id = 2
	t.taxRate = 0.25
	fmt.Printf("%+v\n", t)

	GetMessage(f)
	GetMessage(t)
}
