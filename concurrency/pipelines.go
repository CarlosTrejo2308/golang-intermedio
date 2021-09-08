package main

// Generator generates a stream of integers
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

// Double doubles the value received from the input channel
func Double(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * 2
	}
	close(out)
}

// Print prints the values received from the input channel
func Print(c <-chan int) {
	for i := range c {
		println(i)
	}
}

// main is the entry point for the program
func main() {
	// Create new channels
	in := make(chan int)
	out := make(chan int)

	// Launch the goroutines
	go Generator(in)
	go Double(in, out)
	Print(out)
}
