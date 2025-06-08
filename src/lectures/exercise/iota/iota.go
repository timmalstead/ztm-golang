//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

// lost ten minutes when i forgot that you do NOT use the = operand to assign for a type alias
type Operation int

const (
	add Operation = iota
	sub
	mul
	div
)

var calculations = map[Operation]func(int, int) int{
	add: func(op1, op2 int) int {
		return op1 + op2
	},
	sub: func(op1, op2 int) int {
		return op1 - op2
	},
	mul: func(op1, op2 int) int {
		return op1 * op2
	},
	div: func(op1, op2 int) int {
		return op1 / op2
	},
}

func (op Operation) calculate(op1, op2 int) int {
	var opToExecute, opExists = calculations[op]
	if opExists {
		return opToExecute(op1, op2)
	}
	panic("invalid operation!")
}

// func (op Operation) calculate(op1, op2 int) int {
// 	if op == add {
// 		return op1 + op2
// 	} else if op == sub {
// 		return op1 - op2
// 	} else if op == mul {
// 		return op1 * op2
// 	} else {
// 		return op1 / op2
// 	}
// }

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
