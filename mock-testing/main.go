package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(1 * time.Second)
	// SELECT * FROM person WHERE dni = dni

	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(1 * time.Second)
	// SELECT * FROM employee WHERE...

	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmp FullTimeEmployee

	e, err := GetEmployeeByID(id)
	if err != nil {
		return ftEmp, err
	}
	ftEmp.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmp, err
	}
	ftEmp.Person = p

	return ftEmp, nil
}
