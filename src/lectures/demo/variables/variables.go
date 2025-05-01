package main

import "fmt"

func main() {
	var myName = "Timothy"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("her name =", name)

	userName := "admin"
	fmt.Println("userName =", userName)

	var sum int
	fmt.Println("the sum is", sum)

	part1, other := 1, 5
	fmt.Println("part =", part1, "other =", other)

	part2, other := 2, 0
	fmt.Println("part2 =", part2, "other =", other)

	sum = part1 + part2
	fmt.Println("the sum now equals", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName =", lessonName, "lessonType =", lessonType)

	word1, word2, _ := "hello", "world", 42
	fmt.Println("word1 =", word1, "word2 =", word2)
}
