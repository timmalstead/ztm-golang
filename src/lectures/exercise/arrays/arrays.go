//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type GroceryItem struct {
	productName string
	productCost int
}

// * Print to the terminal:
//   - The last item on the list
//   - The total number of items
//   - The total cost of the items
func findListInformation(list [4]GroceryItem) {
	var (
		// lastItem                             string
		totalNumberOfItems, totalCostOfItems int
	)

	for i := 0; i < len(list); i++ {
		var currentListItem = list[i]

		if currentListItem.productName != "" {
			totalNumberOfItems++
			totalCostOfItems += currentListItem.productCost
			// lastItem = currentListItem.productName
		}
	}

	// fmt.Println("Last item on list:", lastItem, ". Total number of items:", totalNumberOfItems, ". Total cost of items:", totalCostOfItems)
	fmt.Println("Last item on list:", list[totalNumberOfItems-1], ". Total number of items:", totalNumberOfItems, ". Total cost of items:", totalCostOfItems)
}

func main() {
	var shoppingList = [4]GroceryItem{
		{productName: "cereal", productCost: 2},
		{productName: "coffee", productCost: 1},
		{productName: "orange juice", productCost: 2},
	}

	// length is the total length of the array, even if there is no element explicitly put there
	fmt.Println(len(shoppingList), shoppingList[len(shoppingList)-1].productName == "")
	findListInformation(shoppingList)

	// Guess you can't add to an array created using the spread operator. Makes sense I guess, because the size has to be known when instantiated
	shoppingList[3] = GroceryItem{"milk", 3}

	findListInformation(shoppingList)
}
