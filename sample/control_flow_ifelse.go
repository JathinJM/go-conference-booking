package main

import (
	"fmt"
	"strings"
)

func control_flow() {
	conferenceName := "Go Conferennce"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	var firstName string
	var secondName string
	var email string
	var userTicket uint
	var bookings []string

	for {
		fmt.Printf("please enter the first name:\n")
		fmt.Scan(&firstName)

		fmt.Printf("please enter the second name:\n")
		fmt.Scan(&secondName)

		fmt.Printf("please enter the email:\n")
		fmt.Scan(&email)

		fmt.Printf("please enter the user ticket:\n")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) > 2 && len(secondName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTicket := userTicket <= remainingTickets && userTicket > 0

		if isValidName && isValidEmail && isValidUserTicket {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+secondName)

			fmt.Printf("==============================\n")

			fmt.Printf("type of conference ticktes is %T, remaining tickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
			fmt.Printf("We have total of  %v conference ticktes and %v remaining tickets\n", conferenceTickets, remainingTickets)

			var firstNames []string
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Fields(booking)[0])
			}

			fmt.Printf("Booking %v\n", bookings)
			fmt.Printf("First names of Booking %v\n", firstNames)
			fmt.Printf("==============================\n")

			if remainingTickets == 0 {
				fmt.Printf("The conference seats are booked out")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Printf("Invalid Email \n")
			}
			if !isValidName {
				fmt.Printf("Invalid firstName or secondName \n")
			}
			if !isValidUserTicket {
				fmt.Printf("Invalid User Ticket \n")
			}

		}

	}

}
