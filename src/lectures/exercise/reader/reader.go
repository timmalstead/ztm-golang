//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var p, f = fmt.Println, fmt.Printf

// this is such a goofy implmentation
func main() {
	var myNewReader = bufio.NewReader(os.Stdin)

	var isEndOfFile bool
	var commandsEntered, blankLinesEntered int
	for !isEndOfFile {
		var input, inputErr = myNewReader.ReadString('\n')
		var trimmedInput = strings.TrimSpace(input)

		if trimmedInput == "" && inputErr != io.EOF {
			blankLinesEntered++
		} else {
			if strings.ToLower(trimmedInput) == "q" {
				p("You have quit the program")
				inputErr = io.EOF
			} else {
				commandsEntered++

				switch trimmedInput {
				case "hello":
					p("Hello there")
				case "bye":
					p("Goodbye")
				default:
					p("Okay!")
				}
			}
		}
		isEndOfFile = inputErr == io.EOF
	}

	f("you ran %v commands and entered %v blank lines", commandsEntered, blankLinesEntered)
}
