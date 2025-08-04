package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// type Files struct {
// 	paths []string
// 	sync.Mutex
// }

// func validateArgs(grepArgs []string) {
// 	if len(grepArgs) != 2 {
// 		log.Fatal("\nmgrep requires two arguments in an exact order\n1) <search_string>: The string to search for.\n2) <search_dir>: The directory in which to search.")
// 	}
// }

// func validateSearchDirectory(searchDir string) {
// 	var directoryError = os.Chdir(searchDir)
// 	if directoryError != nil {
// 		log.Fatal(directoryError)
// 	}
// 	fmt.Printf("%v is a valid directory.\n", searchDir)
// }

// // I guess function literals aren't able to recurse in go? They seem to need to be in the global scope like this

// func recurseDirectories(dir string, filesChannel chan string) {
// 	var entries, _ = os.ReadDir(dir)

// 	for _, entry := range entries {
// 		var entryName = fmt.Sprintf("%v/%v", dir, entry.Name())
// 		if entry.IsDir() {
// 			recurseDirectories(entryName, filesChannel)
// 		} else if entryName != "" {
// 			filesChannel <- entryName
// 		}
// 	}
// }

// func getFiles(workingDir string) (chan string, context.CancelFunc) {
// 	var filesChannel = make(chan string)

// 	var initialContext = context.Background()
// 	var fullContext, cancelContext = context.WithCancel(initialContext)

// 	go func() {
// 		recurseDirectories(workingDir, filesChannel)
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case <-fullContext.Done():
// 				fmt.Println("closing channel")
// 				close(filesChannel)
// 			}
// 		}
// 	}()

// 	return filesChannel, cancelContext
// }

// func grepFiles(searchString string, filePaths chan string, ctxCancel context.CancelFunc) {
// 	for {
// 		var path = <-filePaths
// 		if path == "" {
// 			ctxCancel()
// 			break
// 		} else {
// 			go func() {
// 				var file, fileOpenErr = os.Open(path)
// 				defer file.Close()

// 				if fileOpenErr != nil {
// 					log.Fatal(fileOpenErr)
// 				}

// 				var fileReader = bufio.NewReader(file)
// 				var currentLine = 1

// 				for {
// 					var line, lineErr = fileReader.ReadString('\n')

// 					var lineHasSearchString = strings.Contains(line, searchString)

// 					if lineHasSearchString {
// 						fmt.Printf("---\nSearch string: %v\nFile path: %v, Line number: %v\nMatching line: %v\n", searchString, path, currentLine, line)
// 					}

// 					currentLine++
// 					if lineErr == io.EOF {
// 						break
// 					}
// 				}
// 			}()
// 		}
// 	}
// }

// func main() {
// 	var grepArgs = os.Args[1:]
// 	validateArgs(grepArgs)

// 	var searchString = grepArgs[0]
// 	var searchDir = grepArgs[1]
// 	validateSearchDirectory(searchDir)

// 	var filePaths, cancelFunc = getFiles(searchDir)

// 	grepFiles(searchString, filePaths, cancelFunc)
// }

//// ------
//// ------
//// ------
//// ------
//// ------

type Files struct {
	paths []string
	sync.Mutex
}

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

func recurseDirectories(dir string, files *Files) {
	var entries, dirErr = os.ReadDir(dir)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(entries))
	for _, entry := range entries {
		go func() {
			defer waitGroup.Done()
			var entryName = fmt.Sprintf("%v/%v", dir, entry.Name())
			if entry.IsDir() {
				recurseDirectories(entryName, files)
			} else {
				files.Lock()
				files.paths = append(files.paths, entryName)
				files.Unlock()
			}
		}()
	}
	waitGroup.Wait()
}

func getFiles(workingDir string) []string {
	var files = Files{}
	recurseDirectories(workingDir, &files)
	return files.paths
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
