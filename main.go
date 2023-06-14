package main

import (
	"fmt"
	"strings"
	// "strconv"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets uint = 50;
	var remainingTickets uint = conferenceTickets;
	var bookings []string

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and are %v still available \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
	
		fmt.Println("Enter Your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter Your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter Your email: ")
		fmt.Scan(&email)
	
		fmt.Println("How many tickets do you want to buy? ")
		fmt.Scan(&userTickets)
	
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
		
		var firstNames = []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		fmt.Printf("These are all our bookings %v \n", bookings)
		fmt.Printf("The first names of bookings are %v \n", firstNames)
	}
}
