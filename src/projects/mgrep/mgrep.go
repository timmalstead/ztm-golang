package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func validateArgs(grepArgs []string) {
	if len(grepArgs) != 2 {
		log.Fatal("\nmgrep requires two arguments in an exact order\n1) <search_string>: The string to search for.\n2) <search_dir>: The directory in which to search.")
	}
}

func validateSearchDirectory(searchDir string) string {
	var directoryError = os.Chdir(searchDir)
	if directoryError != nil {
		log.Fatal(directoryError)
	}

	var absolutePathDir, pathErr = os.Getwd()
	if pathErr != nil {
		log.Fatal(pathErr)
	}

	fmt.Printf("%v is a valid directory\n", absolutePathDir)
	return absolutePathDir
}

func recurseDirectories(dir string, files *[]string) {
	var entries, dirErr = os.ReadDir(dir)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(entries))
	for _, entry := range entries {
		go func() {
			defer waitGroup.Done()
			var entryName = filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				recurseDirectories(entryName, files)
			} else {
				*files = append(*files, entryName)
			}
		}()
	}
	waitGroup.Wait()
}

func getFiles(workingDir string) []string {
	var files = []string{}
	recurseDirectories(workingDir, &files)
	return files
}

func grepFiles(searchString string, filePaths []string) bool {
	var matchFound = false

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(filePaths))

	for _, path := range filePaths {
		go func() {
			defer waitGroup.Done()
			var file, fileOpenErr = os.Open(path)
			defer file.Close()

			if fileOpenErr != nil {
				log.Fatal(fileOpenErr)
			}

			var fileReader = bufio.NewReader(file)
			var currentLine = 1

			for {
				var line, lineErr = fileReader.ReadString('\n')

				var lineHasSearchString = strings.Contains(line, searchString)

				if lineHasSearchString {
					matchFound = true
					fmt.Printf("---\nSearch string: %v\nFile path: %v, Line number: %v\nMatching line: %v\n", searchString, path, currentLine, line)
				}

				currentLine++
				if lineErr == io.EOF {
					break
				}
			}
		}()
	}
	waitGroup.Wait()

	return matchFound
}

func main() {
	var grepArgs = os.Args[1:]
	validateArgs(grepArgs)

	var searchDir = grepArgs[1]
	var absolutePath = validateSearchDirectory(searchDir)

	var filePaths = getFiles(absolutePath)

	var searchString = grepArgs[0]
	var matchesFound = grepFiles(searchString, filePaths)

	if !matchesFound {
		fmt.Printf("No matches for %v were found in %v or any of its subdirectories\n", searchString, absolutePath)
	}
}

////////////////
////////////////
////////////////
////////////////

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"sync"
// )

// const BUFFER_SIZE = 100

// func validateArgs(grepArgs []string) {
// 	if len(grepArgs) != 2 {
// 		log.Fatal("\nmgrep requires two arguments in an exact order\n1) <search_string>: The string to search for.\n2) <search_dir>: The directory in which to search.")
// 	}
// }

// func validateSearchDirectory(searchDir string) string {
// 	var directoryError = os.Chdir(searchDir)
// 	if directoryError != nil {
// 		log.Fatal(directoryError)
// 	}

// 	var absolutePathDir, pathErr = os.Getwd()
// 	if pathErr != nil {
// 		log.Fatal(pathErr)
// 	}

// 	fmt.Printf("%v is a valid directory\n", absolutePathDir)
// 	return absolutePathDir
// }

// func recurseDirectories(dir string, filesChannel chan string) {
// 	var entries, dirErr = os.ReadDir(dir)
// 	if dirErr != nil {
// 		log.Fatal(dirErr)
// 	}

// 	var waitGroup sync.WaitGroup
// 	waitGroup.Add(1)
// 	go func() {
// 		defer waitGroup.Done()
// 		for _, entry := range entries {
// 			var entryName = filepath.Join(dir, entry.Name())
// 			if entry.IsDir() {
// 				recurseDirectories(entryName, filesChannel)
// 			} else {
// 				filesChannel <- entryName
// 			}
// 		}
// 	}()
// 	waitGroup.Wait()
// }

// func getFiles(dir string, filesChannel chan string) {
// 	var waitGroup sync.WaitGroup
// 	waitGroup.Add(1)
// 	go func() {
// 		defer waitGroup.Done()
// 		recurseDirectories(dir, filesChannel)
// 	}()
// 	waitGroup.Wait()
// 	close(filesChannel)
// }

// func grepFiles(searchString string, filesChannel, matchesChannel chan string) {
// 	var waitGroup sync.WaitGroup

// 	for path := range filesChannel {
// 		go func() {
// 			defer waitGroup.Done()
// 			waitGroup.Add(1)

// 			var file, fileOpenErr = os.Open(path)
// 			defer file.Close()

// 			if fileOpenErr != nil {
// 				log.Fatal(fileOpenErr)
// 			}

// 			var fileReader = bufio.NewReader(file)
// 			var currentLine = 1
// 			for {
// 				var line, lineErr = fileReader.ReadString('\n')

// 				var lineHasSearchString = strings.Contains(line, searchString)

// 				if lineHasSearchString {
// 					var notification = fmt.Sprintf("---\nSearch string: %v\nFile path: %v, Line number: %v\nMatching line: %v\n", searchString, path, currentLine, line)
// 					matchesChannel <- notification
// 				}

// 				currentLine++
// 				if lineErr == io.EOF {
// 					break
// 				}
// 			}
// 		}()
// 	}
// 	waitGroup.Wait()
// 	close(matchesChannel)
// }

// func announceResults(searchString, absolutePath string, matchesChannel chan string) {
// 	var matchesFound = false

// 	for match := range matchesChannel {
// 		matchesFound = true
// 		fmt.Println(match)
// 	}

// 	if !matchesFound {
// 		fmt.Printf("No matches for %v were found in %v or any of its subdirectories\n", searchString, absolutePath)
// 	}
// }

// func main() {
// 	var grepArgs = os.Args[1:]
// 	validateArgs(grepArgs)

// 	var searchString = grepArgs[0]
// 	var searchDir = grepArgs[1]
// 	var absolutePath = validateSearchDirectory(searchDir)

// 	var filesChannel = make(chan string, BUFFER_SIZE)
// 	getFiles(absolutePath, filesChannel)

// 	var matchesChannel = make(chan string, BUFFER_SIZE)
// 	grepFiles(searchString, filesChannel, matchesChannel)

// 	announceResults(searchString, absolutePath, matchesChannel)
// }
