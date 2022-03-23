package main

import (
	"booking-go/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName  string
	secondName string
	email      string
	userTicket uint
}

const conferenceTickets int = 50

var conferenceName string = "Go Conferennce"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

func main() {
	greetUser()
	firstName, secondName, email, userTicket := readUserInput()
	isValidName, isValidEmail, isValidUserTicket := helper.ValidateUser(firstName, secondName, email, userTicket, remainingTickets)

	if isValidName && isValidEmail && isValidUserTicket {
		remainingTickets, bookings = getBookings(userTicket, firstName, secondName, email)
		displayConferenceAndBookingDetails()
		fmt.Printf("First names of Booking %v\n", getFirstNames())

		wg.Add(1)
		go generateAndSendTicket(userTicket, firstName, secondName, email)

		if remainingTickets == 0 {
			fmt.Printf("The conference seats are booked out")
		}
	} else {
		printErrorMessage(isValidEmail, isValidName, isValidUserTicket)

	}
	wg.Wait()
}

func getBookings(userTicket uint,
	firstName string, secondName string, email string) (
	uint,
	[]UserData) {
	remainingTickets = remainingTickets - userTicket

	var userData = UserData{
		firstName:  firstName,
		secondName: secondName,
		email:      email,
		userTicket: userTicket,
	}

	bookings = append(bookings, userData)
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
	fmt.Printf("===================\n")
	fmt.Printf("Booking System \n")
	fmt.Printf("Welcome to %s\n", conferenceName)
	fmt.Printf("===================\n")
}

func displayConferenceAndBookingDetails() {
	fmt.Printf("type of conference ticktes is %T, remaining tickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of  %v conference ticktes and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Printf("Bookings List %v\n", bookings)
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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func generateAndSendTicket(userTicket uint,
	firstName string, secondName string, email string) {

	time.Sleep(50 * time.Second)
	fmt.Printf("##################\n")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, secondName)
	fmt.Printf("Sending ticket :\n  %v \n to Email address %v \n", ticket, email)
	fmt.Printf("##################\n")
	wg.Done()
}
