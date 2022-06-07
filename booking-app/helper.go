package main 

import "strings"

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
