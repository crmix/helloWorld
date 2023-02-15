package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 30
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	for {
		var userName string
		var userTickets uint
		var lastName string
		fmt.Print("Enter your first name:")
		fmt.Scan(&userName)

		fmt.Print("Enter last name: ")
		fmt.Scan(&lastName)

		fmt.Print("enter number of ticket: ")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("sooorry now %v tickets are remaining, so you cannot book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName+" "+lastName)

		fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
		fmt.Printf("now %v tickets are available\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break

		}
	}
}
