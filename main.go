package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var conferenceName = "Go Conf"
var confTickets = 50
var confRemTickets = 50
var bookings = make([]User, 0)

// var bookings = make([]map[string]string, 0)

func main() {

	greetUser()

	// Array
	// var bookings [50]string

	// Slice

	for {

		firstName, lastName, email, userTickets := getUserInput()
		// isValidCity := city == "Pune" || city == "Mumbai"

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, confRemTickets)

		if !isValidName && !isValidEmail && !isValidTicketNumber {

			if !isValidName {
				fmt.Printf("FirstName and LastName length should be greater than 2.\n")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email address.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Tickets count should be greater than zero and less than %v", confRemTickets)
			}
			// break // breaks loop
			continue // breaks current itiration
		}

		bookTickets(userTickets, firstName, lastName, email)
		go sendTickets(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("The firstName of bookings are: %v\n", firstNames)

		if confRemTickets == 0 {
			fmt.Println("Conferenece is sold out. come next year.")
			break
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, confRemTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Println(firstNames)

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// var city string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {

	confRemTickets -= userTickets

	// Array
	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("The Whole Array: %v\n", bookings)
	// fmt.Printf("The first valye: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array Size: %v\n", len(bookings))

	// Slice
	// var user = make(map[string]string)

	// user["firstName"] = firstName
	// user["lastName"] = lastName
	// user["email"] = email
	// user["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)
	// fmt.Printf("The Whole Slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice Size: %v\n", len(bookings))
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are not available for %v\n", confRemTickets, conferenceName)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending Ticket %v to email address %v", ticket, email)
	fmt.Println("#####################")
}
