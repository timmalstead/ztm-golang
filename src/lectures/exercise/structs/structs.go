//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// * Create a rectangle structure containing a length and width field
type Rectangle struct {
	Length int
	Width  int
}

//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

func findArea(rec Rectangle) int {
	return rec.Length * rec.Width
}

func findPerimeter(rec Rectangle) int {
	return (rec.Length + rec.Width) * 2
}

func main() {
	//* Using functions, calculate the area and perimeter of a rectangle,
	//  - Print the results to the terminal
	//  - The functions must use the rectangle structure as the function parameter
	var rectangle = Rectangle{Length: 2, Width: 2}
	var initRecArea, initRecPerimeter = findArea(rectangle), findPerimeter(rectangle)

	fmt.Println("starting rectangle area", initRecArea, "starting rectangle perimeter", initRecPerimeter)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	rectangle.Length *= 2
	rectangle.Width *= 2

	var modifiedRecArea, modifiedRecPerimeter = findArea(rectangle), findPerimeter(rectangle)

	fmt.Println("modified rectangle area", modifiedRecArea, "modified rectangle perimeter", modifiedRecPerimeter)
}
