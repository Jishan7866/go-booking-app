package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	conferenceTickets := 50
	var remainingTickets uint = 50
	var bookings = []string{}
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for remainingTickets > 0 && len(bookings) < 50 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {
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

			if remainingTickets == 0 {
				fmt.Println("Our confrence is booked out. come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			} else if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			} else if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			// fmt.Println("Your input data is invalid", "try again !-")
			// continue
		}

	}
	city := "London"
	switch city {
	case "New York":
		///
	case "Singapore":
		///
	case "Berlin":
		///
	case "London":
		//
	case "Mexico":
		//
	case "Hong Kong":
		//
	default:
		fmt.Println("No valid city selected")
	}
}
