package helper

import (
	"strings"
)

func ValidateUser(firstName string,
	secondName string,
	email string,
	userTicket uint,
	remainingTickets uint) (bool,
	bool,
	bool) {
	isValidName := len(firstName) > 2 && len(secondName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTicket <= remainingTickets && userTicket > 0
	return isValidName, isValidEmail, isValidUserTicket
}
