//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

var p, f = fmt.Println, fmt.Printf

type Vehicle struct {
	model string
}

type Motorcycle Vehicle
type Car Vehicle
type Truck Vehicle

type Lifter interface {
	Lift()
}

func (v Motorcycle) Lift() {
	f("small lift for motorcycle %v\n", v.model)
}

func (v Car) Lift() {
	f("medium lift for car %v\n", v.model)
}

func (v Truck) Lift() {
	f("large lift for truck %v\n", v.model)
}

func directLifters(liftables []Lifter) {
	p("liftin' all we can")

	for _, vehicle := range liftables {
		vehicle.Lift()
	}
}

type IntAlias int
type StrAlias string

type AliasChecker interface {
	CheckType()
}

func (i IntAlias) CheckType() {
	f("that's an int: %v\n", i)
}

func (s StrAlias) CheckType() {
	f("that's a string: %v\n", s)
}

func typeChecker(types []AliasChecker) {
	for _, typeAlias := range types {
		typeAlias.CheckType()
		var _, isInteger = typeAlias.(IntAlias)
		var _, isIntegerPointer = typeAlias.(*IntAlias)

		if isInteger {
			p("I told you that was an integer!")
		}
		if isIntegerPointer {
			p("I told you that was an integer pointer!")
		}

		var _, isString = typeAlias.(StrAlias)
		var _, isStringPointer = typeAlias.(*StrAlias)

		if isString {
			p("I told you that was a string!")
		}
		if isStringPointer {
			p("I told you that was a string pointer!")
		}

	}
}

func main() {
	var shark = Motorcycle{"Swifty Shark"}
	var cuda = Car{"Cunning 'Cuda"}
	var hulk = Truck{"Happy Hulk"}

	// use pointers or copies
	directLifters([]Lifter{&shark, &cuda, hulk})

	var intExample, strExample = IntAlias(5), StrAlias("hello")

	typeChecker([]AliasChecker{&intExample, strExample})
}
