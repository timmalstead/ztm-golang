package main

import "fmt"

var p = fmt.Println

type Space struct {
	isOccupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].isOccupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].isOccupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].isOccupied = false
}

func main() {
	var lot = ParkingLot{spaces: make([]Space, 10)}
	p(lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	p(lot)
	lot.vacateSpace(1)
	p(lot)
}
