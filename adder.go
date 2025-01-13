// Package adder is a simple package for adding numbers
package adder

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add allows to add two integers and returns the result
func Add[T Number](a, b T) T {
	return a + b
}
