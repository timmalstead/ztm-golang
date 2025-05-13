package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap!")
	case p < 10:
		fmt.Println("less cheap!")
	default:
		fmt.Println("expensive")
	}

	var ticket int = Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("Best seating")
	default:
		fmt.Println("Other seating")
	}
}
