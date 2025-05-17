package main

import "fmt"

type Room struct {
	name      string
	isCleaned bool
}

// the argument of the array below has to be typed with the length of the array, it cannot be spread like creating an array can
func checkIfClean(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		// good practice in go to cache the working element in a variable to not run into issues with concurrency
		// do this EVEN if you need to work with the iterator value itself. if you need to work with the integer do something like `var iterator = i`
		var currentRoom Room = rooms[i]
		if currentRoom.isCleaned {
			fmt.Println("That room is clean")
		} else {
			fmt.Println("That room is dirty")
		}
	}
}

func main() {
	// all of these will be created with isCleaned as false, as that is the default value for booleans
	var roomsToClean = [...]Room{
		{name: "Antechamber"},
		{name: "Foyer"},
		{name: "Solarium"},
		{name: "Veranda"},
	}
	checkIfClean(roomsToClean)

	fmt.Println("\ngetting to it...\n")

	roomsToClean[0].isCleaned = true
	roomsToClean[2].isCleaned = true

	// since this creates a COPY of the data, changing foyer.isCleaned will not effect the value of the array element it came from
	var foyer = roomsToClean[1]
	foyer.isCleaned = true

	checkIfClean(roomsToClean)

	fmt.Println("\none more time\n")

	roomsToClean[1].isCleaned = true

	checkIfClean(roomsToClean)
}
