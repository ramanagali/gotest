package main

import "fmt"

func main() {
	fmt.Println("Go package")

	fmt.Println(Calculate(5))
}

// Calculate returns x + 2
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
