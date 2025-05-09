//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.

func sayMyName(name string) string {
	return "Hello " + name + "!"
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()

func saySomething() string {
	return "Jet fuel can't melt steel beams"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func addThreeNumbers(int1, int2, int3 int) int {
	return int1 + int2 + int3
}

//* Write a function that returns any number

func giveMeFive() int {
	return 5
}

func returnThreeNumbers(int1, int2, int3 int) (int, int, int) {
	return int1, int2, int3
}

//* Write a function that returns any two numbers

func returnTwoNumbers(int1, int2 int) (int, int) {
	return int1, int2
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	fmt.Println(sayMyName("Tim"))
	fmt.Println(saySomething())
	fmt.Println(addThreeNumbers(1, 2, 3))

	// ignoring third number
	var num1, num2, _ int = returnThreeNumbers(4, 6, 9)
	fmt.Println(num1, num2)

	fmt.Println(giveMeFive())

	var num3, num4 int = returnTwoNumbers(4, 6)
	var sum int = num3 + num4 + giveMeFive()

	fmt.Println(sum)

}
