package main

import "fmt"

// Funciones variadicas
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Retornos con nombres
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4

	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	printNames("Juan", "Pedro", "Maria")
	fmt.Println(getValues(2))
}
