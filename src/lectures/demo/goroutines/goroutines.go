package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	var capIt = func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper((r)))
		fmt.Printf("%c done!\n", r)
	}

	for _, letter := range data {
		go capIt(letter)
	}

	// for i := 0; i < len(data); i++ {
	// 	go capIt(data[i])
	// }

	fmt.Println("Capitalized 1:", capitalized)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Capitalized 2: %c\n", capitalized)
}
