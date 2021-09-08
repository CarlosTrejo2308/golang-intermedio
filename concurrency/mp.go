package main

import (
	"fmt"
	"time"
)

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	/* fmt.Println("Waiting for the first result")
	fmt.Println(<-c1)
	fmt.Println("Waiting for the second result")
	fmt.Println(<-c2) */

	for i := 0; i < 2; i++ {
		select {
		case res := <-c1:
			fmt.Println("Received", res, "from c1")
		case res := <-c2:
			fmt.Println("Received", res, "from c2")
		}
	}

}
