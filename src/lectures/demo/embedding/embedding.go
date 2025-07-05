package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyer interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyer
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	var spamMail = SpamMail{amount: 40000}
	automate(&spamMail)

	// below code will not compile because the argument for the automate function, a pointer to the ToxicWaste struct, points to a type that only have the Shipper interface implemented, and not the conveyer
	// This is a fine way to do this, I guess, but it does seem to go through several unneeded layers of abstraction
	// an easier and more understandable way to do this would be to determine its type in a function
	// defnitely think I will be more into a functional style of go than implementing and embedding interfaces
	// seems a lot like oop and abstract classes by another name, and all of that is just silly
	// var toxicWaste = ToxicWaste{weight: 300}
	// automate(&toxicWaste)
}
