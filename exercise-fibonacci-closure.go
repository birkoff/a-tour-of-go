package main

import "fmt"

// Returns a function that returns successive Fibonacci numbers
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	
	// Print the first 10 Fibonacci numbers
	n := 10
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}
