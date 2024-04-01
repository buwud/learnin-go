package main

import "fmt" //io package

func main() {
	//syntactic sugar -> describe a feauture that lets you do smth more easily
	conferenceName := "Go Conference" //string, create var w this syntax but cannot define type explicitly
	const conferenceTickets int = 50  //cannot change like js
	var remainingTickets uint = 50    //int
	//uint -> positive whole numbers

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	//%T used to print data-type

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string //go doesnt know its type because its not assigned initially, have to add type
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Print(("Enter your first name: "))
	fmt.Scan(&firstName)

	fmt.Print(("Enter your last name: "))
	fmt.Scan(&lastName)

	fmt.Print(("Enter your email address: "))
	fmt.Scan(&email)

	fmt.Print(("Enter number of tickets: "))
	fmt.Scan(&userTickets)

	userTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
