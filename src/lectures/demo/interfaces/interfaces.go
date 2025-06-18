package main

import "fmt"

var p = fmt.Println

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	p("cook that chicken")
}

func (c Salad) PrepareDish() {
	p("chop salad")
	p("add dressing")
}

func prepareDishes(dishes []Preparer) {
	p("preparing dishes")
	for _, dish := range dishes {
		fmt.Printf("Dish: %v\n", dish)

		// okay, can check for different types here, but I'm not exactly sure what I'm doing
		var _, isChicken = dish.(Chicken)
		p(isChicken)

		dish.PrepareDish()
	}
}

func main() {
	var chickenDish Chicken = "Chicken Pot Pie"
	var saladDish Salad = "Endive Salad"

	prepareDishes([]Preparer{chickenDish, saladDish})

}
