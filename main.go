package main

import (
    "fmt"
	// "strconv"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets uint = 50;
	var remainingTickets uint = conferenceTickets;

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and are %v still available \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")

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

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
