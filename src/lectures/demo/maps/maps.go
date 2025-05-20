package main

import "fmt"

func main() {
	var shoppingList = make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)
	fmt.Println("need", shoppingList["eggs"], "eggs")

	var cereal, cerealExists = shoppingList["cereal"]

	if !cerealExists {
		fmt.Println("ain't no cereal")
	} else {
		fmt.Println("cereal is there", cereal)
	}

	var totalItems int
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println(totalItems)

}
