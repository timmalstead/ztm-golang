package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	IsBoarded    bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	var casey = Passenger{Name: "Casey", IsBoarded: false, TicketNumber: 1}

	var (
		bill = Passenger{"Bill", 2, false}
		ella = Passenger{"Ella", 3, false}
	)
	fmt.Println(casey, bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4

	fmt.Println(heidi)

	casey.IsBoarded = true
	bill.IsBoarded = true

	if bill.IsBoarded {
		fmt.Println("Bill has done it!")
	} else {
		fmt.Println("Bill can't get on!")
	}

	if casey.IsBoarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.IsBoarded = true

	var crossTownExpress = Bus{FrontSeat: heidi}

	fmt.Println(crossTownExpress.FrontSeat.Name, "in the front seat, party in the back")
}
