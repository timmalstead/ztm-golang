package main

import "fmt"

func main() {
	fmt.Println("ranges")
	var exampleSlice = []string{"hello", "world", "!"}

	// range returns TWO things, the index and the value
	for i, element := range exampleSlice {
		fmt.Println(i, element)
		for _, char := range element {
			// convert numeric rune value to character
			var charRuneConvertedToGlyph = fmt.Sprintf("%q", char)
			fmt.Print(charRuneConvertedToGlyph)
			// fmt.Printf("%q", char)
		}
		fmt.Println()
	}
}
