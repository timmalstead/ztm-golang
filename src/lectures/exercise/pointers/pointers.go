//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

// * Create functions to activate and deactivate security tags using pointers

func activateTag(tag *bool) {
	*tag = Active
}

func deactivateTag(tag *bool) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice

func checkout(tags []*bool) {
	for _, tag := range tags {
		deactivateTag(tag)
	}
}

func main() {
	fmt.Println("Pointers exercise")

	//  - Create at least 4 items, all with active security tags
	var timTag, sophiaTag, bobTag, billTag bool

	activateTag(&timTag)
	activateTag(&sophiaTag)
	activateTag(&bobTag)
	activateTag(&billTag)

	//  - Store them in a slice or array
	var tagHolder = []bool{timTag, sophiaTag, bobTag, billTag}

	fmt.Println("original tags:", timTag, sophiaTag, bobTag, billTag)
	fmt.Println("tags copied into slice:", tagHolder)

	//  - Deactivate any one security tag in the array/slice

	deactivateTag(&tagHolder[2])
	deactivateTag(&tagHolder[3])

	fmt.Println("original tags:", timTag, sophiaTag, bobTag, billTag)
	fmt.Println("tags copied into slice:", tagHolder)
	//  - Call the checkout() function to deactivate all tags
	checkout([]*bool{&tagHolder[0], &tagHolder[1], &tagHolder[2], &tagHolder[3]})
	//  - Print out the array/slice after each change

	fmt.Println("original tags:", timTag, sophiaTag, bobTag, billTag)
	fmt.Println("tags copied into slice:", tagHolder)

}
