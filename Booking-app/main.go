package main

import (
	"fmt"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)
var remainingTickets uint = 50

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNum := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNum {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings first names: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Sorry, all tickets are sold out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Sorry, your first or last name is too short. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Sorry, your email address is invalid. Please try again.")
			}
			if !isValidTicketNum {
				fmt.Println("Sorry, your ticket number is invalid. Please try again.")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to the %v Booking App!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v remaining.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
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

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(
	userTickets uint,
	firstName string,
	lastName string,
	email string,
) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you, %v %v, for your purchase of %v tickets to %v. An email will be sent to %v.\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining.", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprint("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n to %v \nto email address%v.\n", ticket, email)
	fmt.Println("#######################")

}
