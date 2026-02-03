package main

import "fmt"

func main() {
	//variable - value can change
	var conferenceName = "Go Conference"
	//constants - value cannot change
	const conferenceTickets = 50
	var remainingTickets = 30

	fmt.Printf("Welcome to %v booking applicatio\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("We have %v tickets! Currenlty %v tickets left available now!!", conferenceTickets, remainingTickets)
}
