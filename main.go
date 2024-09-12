package main

import (
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference 2024" // only for var variable declaration
const conferenceTicket int = 50

var remainingTickets uint = 50              // positive integer
var bookings = make([]map[string]string, 0) // slice of map

func main() {

	greetUser()

	for {
		// Ask user to enter the input
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {

			bookTicket(firstName, lastName, userTickets, email)
			fmt.Printf("List of bookings: %v\n", bookings)

			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out")
				break
			}
		} else if userTickets == remainingTickets {
			// do something else
		} else {
			if !isValidName {
				fmt.Printf("Your name is invalid. Try again!\n")
			}
			if !isValidEmail {
				fmt.Printf("Your email is invalid. Try again!\n")
			}
			if !isValidTickets {
				fmt.Printf("Your ticket number is invalid. Try again!")
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get you ticket  here to attend %v\n", conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets = remainingTickets - userTickets

	//Create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // convert uint to string

	bookings = append(bookings, userData)

	fmt.Printf("The whole slice are: %v\n", bookings)
	fmt.Printf("The first value of bookings is: %v\n", (bookings)[0])
	fmt.Printf("The slice length is: %v\n", len(bookings))

	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
