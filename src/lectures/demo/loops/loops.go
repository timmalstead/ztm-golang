package main

import "fmt"

func main() {
	var sum int = 0
	fmt.Println("sum:", sum)

	// var i int = 1
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("sum:", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("sum:", sum)
	}
}
