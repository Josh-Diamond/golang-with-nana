package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	// array = [fixed length]datatype value
	//var bookings = [50]string{"Nana", "Remy", "Josh"}

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var rsvp uint = 1
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("%v parties are attending this event\n", rsvp)
}
