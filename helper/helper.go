package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTickets int, confRemTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < confRemTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
