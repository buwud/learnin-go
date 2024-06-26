package main

import (
	"fmt" //io package
	"learning-go/helper"
	"sync"
	"time"
)

// Package Level Variables - Clean Code
const conferenceTickets int = 50     //cannot change like js
var conferenceName = "Go Conference" //string, create var w this syntax but cannot define type explicitly
var remainingTickets uint = 50       //int
// uint -> positive whole numbers
var bookings = make([]UserData, 0) //create empty list of maps, add initial size and it will increase

// STRUCT
// type --> creates new type with the name that specified
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	//syntactic sugar -> describe a feauture that lets you do smth more easily
	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	//%T used to print data-type

	//only have for loop FOR loop c:
	for {
		firstName, lastName, email, userTickets := getUserInput()
		//isValidCity := city == "Singapore" || city == "London"
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) //1 goroutine
			go sendTicket(userTickets, firstName, lastName, email)
			//go starts a new goroutine, creates new thread

			//call func
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

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
	wg.Wait() //waits for all threads to be done before application exits
}

func greetUsers() { //only executed when called
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	//blank identifier (_) -> to ignore a variable you dont want to use
	for _, booking := range bookings {
		//var names = strings.Fields(booking) //splits the string with white space as seperator,
		//returns a slice with the split element
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	//fmt.Printf("The first names of bookings are: %v \n", firstNames)
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName
	//create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// CONCURRENCY
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
