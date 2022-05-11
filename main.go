package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// Array
	var bookings [50]string
	// Slice
	var bookings_slice []string

	// Syntactic Sugar in Go language
	// smooth and more easily way to define and initialize variables
	// make language "sweeter" for human use.

	dabo := "Dabo"
	fmt.Printf("My Name is: %v \n", dabo)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking App!")
	fmt.Printf("Welcome to %v booking App!\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// loop to keep asking after one booking is done.
	for remainingTickets > 0 && len(bookings_slice) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		//push
		// ask user for their data
		fmt.Println("Please Enter your first name: ")
		fmt.Scan(&firstName)
		//checking that first name more than or equal to 3 characters
		for len(firstName) < 3 {
			fmt.Println("Please Enter a Valid first name with more than 3 characters")
			fmt.Scan(&firstName)
		}

		fmt.Println("Please Enter your last name: ")
		fmt.Scan(&lastName)
		//checking that last name more than or equal to 3 characters
		for len(lastName) < 3 {
			fmt.Println("Please Enter a Valid last name with more than 3 characters")
			fmt.Scan(&lastName)
		}

		fmt.Println("Please Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Please Enter number of tickets: ")
		fmt.Scan(&userTickets)
		//only takes input while the desired count of tickets is smaller than remainingTickets
		for remainingTickets < userTickets {
			fmt.Printf("Please Enter a number of tickets less than or equal to %v because these are the only remaining tickets\n", remainingTickets)
			fmt.Scan(&userTickets)
		}

		// add data to our array
		bookings[0] = firstName + " " + lastName
		fmt.Printf("The Whole Array: %v\n", bookings)
		fmt.Printf("The first Value in the Array: %v\n", bookings[0])
		fmt.Printf("The Array Type: %T\n", bookings)
		fmt.Printf("The Array Length: %v\n", len(bookings))

		// add data to our dynamic array (Slice)
		bookings_slice = append(bookings_slice, firstName+" "+lastName)
		//adding a value to slice is different from adding to array,
		//but retreiving data from slice same retreiving from array
		fmt.Printf("The Whole Slice: %v\n", bookings_slice)
		fmt.Printf("The first Value in the Slice: %v\n", bookings_slice[0])
		fmt.Printf("The Slice Type: %T\n", bookings_slice)
		fmt.Printf("The Slice length: %v\n", len(bookings_slice))

		firstNames := []string{}
		//foreach in GoLang
		//I've used _ instead of index variable to tell go that i explicitly don't want that variable
		for _, booking := range bookings_slice {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all of the firstNames of bookings: %v\n", firstNames)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

		//If condition of a boolean variable
		//to break from the loop if true
		//Checking if the remainingTickets variable is equal to 0
		//End the program
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Printf("All the tickets of our conference are sold out\n")
			break
		}
	}
}
