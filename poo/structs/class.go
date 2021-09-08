package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// Reciver functions
func (e Employee) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", e.id, e.name)
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

func main() {
	e := Employee{}
	fmt.Printf("%+v\n", e)

	e.id = 1
	e.name = "Jane"
	fmt.Printf("%+v\n", e)

	e.SetId(5)
	e.SetName("John")
	fmt.Printf("%+v\n", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
