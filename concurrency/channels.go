package main

import "fmt"

func main() {
	// Unbuffered channel
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 1

	// Buffered channel
	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)

	ch2 <- 3
	fmt.Println(<-ch2)

}
