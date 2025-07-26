//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "time"
)

func main() {
	var files = []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	var grandTotal, filesProcessed int

	var incrementFilesProcessed = func() {
		filesProcessed++
	}

	var processNumberFile = func(filePath string) {
		// looks like dlv as controlled by vs code always runs the debugged script in the directory it is in, which means if you are opening files you need to be in that directory as well.
		var expandedPath = fmt.Sprintf("./exercise/goroutines/%v", filePath)
		var file, fileOpenErr = os.Open(expandedPath)
		defer file.Close()
		// this way it will increment the file processing counter whether it is successful or not
		defer incrementFilesProcessed()

		if fileOpenErr != nil {
			fmt.Println("Error opening file: ", fileOpenErr)
		} else {
			var fileTotal int

			var fileReader = bufio.NewReader(file)
			for {
				var line, lineErr = fileReader.ReadString('\n')

				if lineErr == io.EOF {
					break
				}

				var trimmedLine = strings.TrimSpace(line)
				var num, convertErr = strconv.Atoi(trimmedLine)

				if convertErr == nil {
					fileTotal += num
				}

			}

			grandTotal += fileTotal
			// filesProcessed++
		}
	}

	// this makes sense as a candidate to be parallelized as we don't care about the order of the calculations, just the total sum
	for _, file := range files {
		go processNumberFile(file)
	}

	fmt.Println(grandTotal)
	// time.Sleep(100 * time.Millisecond)
	//* The grand total for the files is 4103109

	// guess you can wait on the parallel call with another loop. makes sense
	// there would probably need to be something else though, like a way to terminate if it goes over a timeout
	// I hope I learn about that in future lessons
	for filesProcessed < len(files) {
		fmt.Print(".")
	}

	fmt.Println("\n", grandTotal, grandTotal == 4103109)
}
