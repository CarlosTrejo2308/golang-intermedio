package main

import "fmt"

// Worker is a function that does the work.
func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		fib := Fibonnaci(j)
		results <- fib
		fmt.Println("worker", id, "finished job", j, "result", fib)
	}
}

// Fibonacci returns the nth number in the Fibonacci sequence.
func Fibonnaci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}

func main() {
	// tasks to do
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 35, 37, 40, 41, 42}

	nWorkers := 5
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// start the workers
	for w := 1; w <= nWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// give the workers jobs
	for _, t := range tasks {
		jobs <- t
	}
	close(jobs)

	// get the results (consume the channel)
	for a := 1; a <= len(tasks); a++ {
		<-results
	}

}
