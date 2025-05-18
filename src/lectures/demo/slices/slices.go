package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println()
	for i := 0; i < len(slice); i++ {
		// remember that we are making a COPY of the element in the slice and this will help with concurrency errors, when we get to that
		var currentElement = slice[i]
		fmt.Println(currentElement)
	}
}

func main() {
	var route = []string{"Grocery Store", "Mall", "Salon"}

	printSlice("route 1", route)

	route = append(route, "Home")

	printSlice("route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[2], "visited")

	route = route[2:]

	printSlice("route 3", route)
}
