package adder

import "golang.org/x/exp/constraints"

// Adder is a simple generic function that adds two numbers.
// It is a demonstration of how to create a Go module.
// This function takes two numbers as input and returns their sum.
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
// 		result2 := adder.Add(2.3, 3.2)
// 		fmt.Println(result2) // Output: 5.5
// 	}
//
type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	return a + b
}
