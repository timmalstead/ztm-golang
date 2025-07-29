//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {
	//* Count the total number of letters in any chosen input
	var totalLetterCount int
	//* The input must be supplied from standard input
	var stdInReader = bufio.NewReader(os.Stdin)

	var separatedStrings = []string{}
	// I don't THINK that there is a way to do this reading in goroutines
	// as I understand it the nature of the reader is getting some bytes, consuming them and then waiting for more to work with
	// so MAYBE we could do it in a goroutine if I knew the sizes of the lines beforehand, but I'm not going to try that
	for {
		var input, inputErr = stdInReader.ReadString('\n')
		if inputErr == io.EOF {
			break
		}
		var trimmedInput = strings.TrimSpace(input)
		var splitInput = strings.Split(trimmedInput, " ")
		separatedStrings = append(separatedStrings, splitInput...)
	}
	//* Input analysis must occur per-word, and each word must be analyzed
	//  within a goroutine
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(separatedStrings))

	var countLetters = func(word string) {
		defer waitGroup.Done()
		totalLetterCount += len(word)
	}

	// didn't really see any need to use a mutex on this one
	for _, word := range separatedStrings {
		go countLetters(word)
	}

	waitGroup.Wait()

	//* When the program finishes, display the total number of letters counted
	// should be 29 letters counted
	fmt.Println("totalLetterCount", totalLetterCount)
}
