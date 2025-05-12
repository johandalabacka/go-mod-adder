package adder
// Adder is a simple function that adds two integers.
// It is a demonstration of how to create a Go module.
// This function takes two integers as input and returns their sum.
// It is a simple example of how to create a Go module and use it in another Go program.
// Example usage:
//
// 	package main
//
// 	import (
// 		"fmt"
// 		"github.com/johandalabacka/go-mod-adder/adder"
// 	)
//
// 	func main() {
// 		result := adder.Add(2, 3)
// 		fmt.Println(result) // Output: 5
// 	}
//
func Add(a, b int) int {
	return a + b
}
