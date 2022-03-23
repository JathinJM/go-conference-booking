package main

import (
	"booking-go/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conferennce"
var remainingTickets uint = 50
var bookings []string

func main() {
	greetUser()
	for {
		firstName, secondName, email, userTicket := readUserInput()
		isValidName, isValidEmail, isValidUserTicket := helper.ValidateUser(firstName, secondName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicket {
			remainingTickets, bookings = getBookings(userTicket, firstName, secondName)
			displayConferenceAndBookingDetails()
			fmt.Printf("First names of Booking %v\n", getFirstNames())

			if remainingTickets == 0 {
				fmt.Printf("The conference seats are booked out")
				break
			}
		} else {
			printErrorMessage(isValidEmail, isValidName, isValidUserTicket)

		}

	}

}

func getBookings(userTicket uint,
	firstName string, secondName string) (
	uint,
	[]string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName+" "+secondName)
	return remainingTickets, bookings
}

func readUserInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var userTicket uint

	fmt.Printf("please enter the first name:\n")
	fmt.Scan(&firstName)

	fmt.Printf("please enter the second name:\n")
	fmt.Scan(&secondName)

	fmt.Printf("please enter the email:\n")
	fmt.Scan(&email)

	fmt.Printf("please enter the user ticket:\n")
	fmt.Scan(&userTicket)
	return firstName, secondName, email, userTicket
}

func greetUser() {

	fmt.Printf("Booking System \n")
	fmt.Printf("Welcome to %s\n", conferenceName)
}

func displayConferenceAndBookingDetails() {
	fmt.Printf("type of conference ticktes is %T, remaining tickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of  %v conference ticktes and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Printf("Booking %v\n", bookings)
}

func printErrorMessage(isValidEmail bool, isValidName bool, isValidUserTicket bool) {
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

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
	return firstNames
}
