package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// Handling errors
	myValue, err := strconv.ParseInt("Nan", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", myValue)
	}

	// Map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	// Slice
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println()

	// Append slice
	s = append(s, 6)

	// Range
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println()

	// Delete from slice
	s = append(s[:2], s[3:]...)

	// Range
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Bob", 20}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	// Pointer
	var p2 *Person
	p2 = &p
	fmt.Println(p2.Name)
	fmt.Println(p2.Age)

	fmt.Println()

	// Go routines
	c := make(chan int)
	go doSomething(c)

	fmt.Println(<-c)

	fmt.Println()

	// More pointers
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(*h)

	*h = 30
	fmt.Println(g)

}

func doSomething(c chan int) {
	fmt.Println("Doing something")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")

	c <- 1
}
