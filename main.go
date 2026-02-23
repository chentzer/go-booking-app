package main

import "fmt"

func main() {
	//variable - value can change
	var conferenceName = "Go Conference"
	//constants - value cannot change
	const conferenceTickets = 50
	var remainingTickets = 30

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking applicatio\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("We have %v tickets! Currenlty %v tickets left available now!!", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	// ask user for thier name

	userName = "Tom"
	userTickets = 2
	fmt.Println("User %v booked %v tickets.\n", userName, userTickets)
}
