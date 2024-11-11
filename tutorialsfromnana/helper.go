package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var validName bool = len(firstName) >= 2 && len(lastName) >= 2
	var validEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var validTickets bool = userTickets > 0 && userTickets <= remainingTickets

	return validName, validEmail, validTickets
}