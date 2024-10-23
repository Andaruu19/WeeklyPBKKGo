package main

import (
	"fmt"
	"strings"
)

func main() { 
	conferenceName := "PBKK"
	const conferenceTickets = 10
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v booking Apps\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("Please enter your first name")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter your email")
		fmt.Scanln(&email)

		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scanln(&userTickets)

		// User Validation
		var validName bool = len(firstName) >= 2 && len(lastName) >= 2
		var validEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
		var validTickets bool = userTickets > 0 && userTickets <= remainingTickets

		if validEmail && validName && validTickets {
			remainingTickets -= userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)

			fmt.Printf("There are %v tickets remaining\n", remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Data type : %T\n", bookings)
			fmt.Printf("Data : %v\n", bookings)
			fmt.Printf("Data length : %v\n", len(bookings))
		} else {
			fmt.Println("There are many invalid inputs")
		}

		if remainingTickets <= 0 {
			fmt.Println("All tickets are sold out")
			break
		}

	}

	// SWITCH CASE
	fmt.Println("Tolong masukkan asal kota")
	var nation string
	fmt.Scan(&nation)

	nation = strings.ToLower(nation)

	switch nation {
	case "Indonesia", "Japan", "Korea":
		fmt.Println("Hello from Asia")
	case "United Kingdom", "Germany", "Italy":
		fmt.Println("Hello from Europe")
	case "Lowokwaru", "Ngantang":
		fmt.Println("Woo wong malang iki")
	default:
		fmt.Println("Where are you from???!?")
	}

}