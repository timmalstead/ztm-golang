//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func showAssemblyLine(line []Part) {
	fmt.Println()
	for i := 0; i < len(line); i++ {
		var currentPart = line[i]
		fmt.Println(currentPart)
	}
}

func main() {
	//  - Create an assembly line having any three parts
	var parts = []Part{"beep", "boop", "blip"}
	showAssemblyLine(parts)
	//  - Add two new parts to the line
	parts = append(parts, "blap", "burs")
	showAssemblyLine(parts)
	//  - Slice the assembly line so it contains only the two new parts
	parts = parts[3:]
	showAssemblyLine(parts)
}
