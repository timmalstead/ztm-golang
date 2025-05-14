//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		//* Print integers 1 to 50, except:
		var valToPrint string

		var divisibleByThree, divisibleByFive = i%3 == 0, i%5 == 0
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		if divisibleByThree && divisibleByFive {
			valToPrint = "FizzBuzz"
			//  - Print "Fizz" if the integer is divisible by 3
		} else if divisibleByThree {
			valToPrint = "Fizz"
			//  - Print "Buzz" if the integer is divisible by 5
		} else if divisibleByFive {
			valToPrint = "Buzz"
		} else {
			valToPrint = fmt.Sprint(i)
		}

		fmt.Println(valToPrint)
	}

	for i := 1; i <= 50; i++ {
		var valToPrint string

		if i%3 == 0 {
			valToPrint += "Fizz"
		}
		if i%5 == 0 {
			valToPrint += "Buzz"
		}
		if valToPrint == "" {
			valToPrint = fmt.Sprint(i)
		}

		fmt.Println(valToPrint)
	}
}
