package main

import "fmt"

func sum(nums ...int) int {
	var finalSum int
	for _, n := range nums {
		finalSum += n
	}
	return finalSum
}

func main() {
	fmt.Println("variadics!")
	// like the spread param in js, variadics are a way to write a function that accepts any number of params
	var (
		slice1 = []int{1, 2, 3}
		slice2 = []int{4, 5, 6}
	)

	// remember, the spread in go is AFTER the value to be spread
	fmt.Println(sum(slice1...))
	fmt.Println(sum(slice2...))

	var compoundSlice = append(slice1, slice2...)
	fmt.Println(sum(compoundSlice...))
}
