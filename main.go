package main

import (
	"fmt" //io package
	"strings"
)

func main() {
	//syntactic sugar -> describe a feauture that lets you do smth more easily
	conferenceName := "Go Conference" //string, create var w this syntax but cannot define type explicitly
	const conferenceTickets int = 50  //cannot change like js
	var remainingTickets uint = 50    //int
	//uint -> positive whole numbers
	var bookings []string //array type

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	//%T used to print data-type

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//only have for loop FOR loop c:
	for {
		var firstName string //go doesnt know its type because its not assigned initially, have to add type
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Print(("Enter your first name: "))
		fmt.Scan(&firstName)

		fmt.Print(("Enter your last name: "))
		fmt.Scan(&lastName)

		fmt.Print(("Enter your email address: "))
		fmt.Scan(&email)

		fmt.Print(("Enter number of tickets: "))
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The first value: %v\n", bookings[0])

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			//blank identifier (_) -> to ignore a variable you dont want to use
			for _, booking := range bookings {
				var names = strings.Fields(booking) //splits the string with white space as seperator,
				//returns a slice with the split element
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			// noTicketRemaining bool = remainingTickets == 0
			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				// end the program
				fmt.Println("Conference is booked out. Come back next year!")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remain, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}
}
