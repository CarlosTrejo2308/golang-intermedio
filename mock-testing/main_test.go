package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "12345678",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "12345678",
					Name: "John Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	// Save the original functions
	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		// Set the mock functions for the test
		test.mockFunc()

		// Call the function to test
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error getting full time employee by id: %v", err)
		}

		// Check if the result is the expected
		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Expected age %v, got %v", test.expectedEmployee.Person.Age, ft.Person.Age)
		}

		if ft.DNI != test.expectedEmployee.DNI {
			t.Errorf("Expected dni %v, got %v", test.expectedEmployee.Person.DNI, ft.Person.DNI)
		}

		if ft.Name != test.expectedEmployee.Name {
			t.Errorf("Expected name %v, got %v", test.expectedEmployee.Person.Name, ft.Person.Name)
		}

		if ft.Employee.Id != test.expectedEmployee.Employee.Id {
			t.Errorf("Expected employee id %v, got %v", test.expectedEmployee.Employee.Id, ft.Employee.Id)
		}

		if ft.Employee.Position != test.expectedEmployee.Employee.Position {
			t.Errorf("Expected employee position %v, got %v", test.expectedEmployee.Employee.Position, ft.Employee.Position)
		}
	}

	// Restore the original functions
	GetEmployeeByID = originalGetEmployeeByID
	GetPersonByDNI = originalGetPersonByDNI
}
