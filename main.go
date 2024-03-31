package main

import "fmt" //io package

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 //cannot change like js
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
