package main

import (
	"fmt"
	"time"
)

// Funciones anónimas
func main() {
	func() {
		println("Hello")
	}()

	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("Finishing function")
		c <- 1
	}()
	fmt.Println(<-c)
}
