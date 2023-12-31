package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		curr = prev + curr
		prev = curr - prev
		return curr - prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
