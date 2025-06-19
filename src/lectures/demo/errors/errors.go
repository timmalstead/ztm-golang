package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		// guess there is a simpler way to do errors now
		// return 0, errors.New(fmt.Sprintf("no element at index %v", index))
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	var stuff = Stuff{values: []int{1}}
	var value, err = stuff.Get(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found that integer", value)
	}
}
