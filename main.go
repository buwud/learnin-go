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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	//%T used to print data-type

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 //validate -> at least 2 chars
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The first value: %v\n", bookings[0])

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//call func
			printFirstNames(bookings)

			// noTicketRemaining bool = remainingTickets == 0
			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				// end the program
				fmt.Println("Conference is booked out. Come back next year!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number if tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid, try again.")
		}
	}

	city := "London"

	switch city {
	case "New York":
	case "Singapore", "Hong Kong":
	case "Berlin", "London":
	case "Mexico":
	default:
		fmt.Print("No valid city selected")
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) { //only executed when called
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func printFirstNames(bookings []string) {
	firstNames := []string{}
	//blank identifier (_) -> to ignore a variable you dont want to use
	for _, booking := range bookings {
		var names = strings.Fields(booking) //splits the string with white space as seperator,
		//returns a slice with the split element
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are: %v \n", firstNames)
}
