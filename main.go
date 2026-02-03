package main

import "fmt"

func main() {
	//variable - value can change
	var conferenceName = "Go Conference"
	//constants - value cannot change
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend!")
	fmt.Println("We have", conferenceTickets, "tickets! Currenlty", remainingTickets, "tickets left available now!!")
}
