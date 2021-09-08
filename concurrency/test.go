package main

import (
	"fmt"
	"time"
)

func doSmth(i int, c chan int) {
	fmt.Printf("Id: %d -> started...\n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Id: %d -> finished...\n", i)
	<-c // frees the space for new routines
}

func main() {
	c := make(chan int, 5) // creates a buffered channel with a capacity of two

	for i := 0; i < 10; i++ {
		c <- 1 // alocate a new "instance" in the free space
		go doSmth(i, c)
	}

	fmt.Println("Waiting for all routines to finish...")
}
