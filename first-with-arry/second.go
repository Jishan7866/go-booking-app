package main

import (
	"fmt"
	"strings"
)

func test() {
	conferenceName := "Go conference"
	conferenceTickets := 50
	var remainingTickets uint = 50
	var bookings = []string{}
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first Name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last Name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email  address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are:%v\n", firstNames)
		noTicketsRemainig := remainingTickets == 0
		if noTicketsRemainig {
			fmt.Println("Our confrence is booked out. come back next year")
			break
		}
	}
}
