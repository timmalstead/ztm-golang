package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	var quiz1, quiz2, quiz3 int = 9, 9, 8
	fmt.Println(quiz1, quiz2, quiz3)

	if quiz1 > quiz2 {
		fmt.Println(("quiz1 wins"))
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 wins")
	} else {
		fmt.Println("scores equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("right on")
	} else {
		fmt.Println("study harder")
	}
}
