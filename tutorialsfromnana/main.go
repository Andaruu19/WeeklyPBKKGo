package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceName = "PBKK"
const conferenceTickets = 10
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	NumberofTickets uint
}

var wg = sync.WaitGroup{}

func main() { 


	greetUsers()

		firstName, lastName, email, userTickets := getUserInput()
		
		// User Validation
		validEmail, validName, validTickets := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


		if validEmail && validName && validTickets {
			bookTicket(userTickets, firstName, lastName, conferenceName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			FirstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", FirstNames)

			fmt.Printf("Data type : %T\n", bookings)
			fmt.Printf("Data : %v\n", bookings)
			fmt.Printf("Data length : %v\n", len(bookings))
		} else {
			fmt.Println("There are invalid inputs")
		}
		if remainingTickets <= 0 {
			fmt.Println("All tickets are sold out")
		}
		wg.Wait()


	// SWITCH CASE
	fmt.Println("Please enter your place of origin")
	var nation string
	fmt.Scan(&nation)

	nation = strings.ToLower(nation)

	switch nation {
	case "indonesia", "japan", "korea":
		fmt.Println("Hello from Asia")
	case "united kingdom", "germany", "italy":
		fmt.Println("Hello from Europe")
	case "lowokwaru", "ngantang":
		fmt.Println("Woo wong malang iki")
	default:
		fmt.Println("Where are you from???!?")
	}


}

func greetUsers() {
	fmt.Printf("Welcome to %v booking Apps\n", conferenceName)
	fmt.Println("Welcome to our conference")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter your email")
		fmt.Scanln(&email)

		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, conferenceName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		NumberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################################################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#####################################################")
	wg.Done()
}