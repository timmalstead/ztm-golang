package main

import "fmt"

// so functions as params pretty much
func add(leftOp, rightOp int) int {
	return leftOp + rightOp
}

func compute(leftOp, rightOp int, op func(leftOp, rightOp int) int) int {
	fmt.Printf("Running computation with %v and %v\n", leftOp, rightOp)
	return op(leftOp, rightOp)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))
	// inline functions work too
	fmt.Println("10 - 2 =", compute(10, 2, func(leftOp, rightOp int) int {
		return leftOp - rightOp
	}))

	// So this works inside the function scope, and I guess you can't have keyword functions inside the function scope.
	var multiply = func(leftOp, rightOp int) int {
		fmt.Println("chills, multiplying")
		return leftOp * rightOp
	}

	fmt.Println("2 x 3 =", compute(2, 3, multiply))
}
