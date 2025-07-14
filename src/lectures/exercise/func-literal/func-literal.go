//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type CharCounterFuncs map[string]func(rune) bool

func findCharCounts(lines []string, funcs CharCounterFuncs) {
	var counter = map[string]int{}
	for _, line := range lines {
		for _, char := range line {
			for funcTitle, counterFunc := range funcs {
				if counterFunc(char) {
					counter[funcTitle]++
					break
				}
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	var lines = []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	var counterFunctions = CharCounterFuncs{
		"letter":      unicode.IsLetter,
		"space":       unicode.IsSpace,
		"digit":       unicode.IsDigit,
		"punctuation": unicode.IsPunct,
	}

	findCharCounts(lines, counterFunctions)
}
