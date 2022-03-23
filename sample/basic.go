package main

import "fmt"

func base() {
	conferenceName := "Go Conferennce"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var firstName string
	var secondName string
	var userTicket uint
	var bookingList [50]string

	var bookings []string

	fmt.Printf("please enter the first name:\n")
	fmt.Scan(&firstName)

	fmt.Printf("please enter the second name:\n")
	fmt.Scan(&secondName)

	fmt.Printf("please enter the user ticket:\n")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	bookingList[0] = firstName + " " + secondName
	bookings = append(bookings, firstName+" "+secondName)

	fmt.Printf("==============================\n")
	fmt.Printf("Booking System")
	fmt.Printf("Welcome to %s\n", conferenceName)
	fmt.Printf("Name is %s\n", firstName)
	fmt.Printf("type of conference ticktes is %T, remaining tickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of  %v conference ticktes and %v remaining tickets\n", conferenceTickets, remainingTickets)

	fmt.Printf("Booking array %v\n", bookingList)
	fmt.Printf("Booking array length %v\n", len(bookingList))
	fmt.Printf("Booking array type %T\n", bookingList)
	fmt.Printf("Booking first element %v\n", bookingList[0])

	fmt.Printf("Booking slice %v\n", bookings)
	fmt.Printf("Booking slice length %v\n", len(bookings))
	fmt.Printf("Booking slice type %T\n", bookings)
	fmt.Printf("Booking slice first element %v\n", bookings[0])
	fmt.Printf("==============================\n")

}
