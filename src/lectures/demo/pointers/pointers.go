package main

import "fmt"

type Counter struct {
	hits int
}

func increment(countner *Counter) {
	// seems good to know. with a REGULAR variable we would need to dereference with *, but using dot notation with a struct, we do not need to do that. I suspect this is because it is also a pointer
	countner.hits++
	fmt.Println("Counter", countner)
}

func replaceStr(old *string, new string, counter *Counter) {
	// remember to dereference!
	*old = new
	increment(counter)
}

// by default, go uses "pass by value" when calling functions, meaning a COPY of information supplied as args. This can get quite unwieldy and it can be difficult to manage state, as not all copies of data may be the same. This can be helped by using pointers, which point to the original data in memory.
func main() {
	fmt.Println("Pointers")

	var counter = Counter{}
	fmt.Println(counter)

	var hello = "Hello"
	var world = "world!"

	fmt.Println(hello, world)

	replaceStr(&world, "there!", &counter)

	fmt.Println(hello, world)
	fmt.Println(counter)

	// making a copy of the data in a new slice. I predict keeping track of what is a copy and what is not will be tricky at first
	var phrase = []string{hello, world}
	fmt.Println(phrase)

	replaceStr(&phrase[1], "Go!", &counter)

	// note that this printing will NOT Have hello go! because we created a COPY when we put these variables into a slice
	fmt.Println(hello, world)
	fmt.Println(phrase)

	fmt.Println(counter)
}
