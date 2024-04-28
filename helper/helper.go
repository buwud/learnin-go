package helper

import "strings"

//capitalizing first letter leads to exporting a variable
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 //validate -> at least 2 chars
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

//3 Levels of scope
//local(within func)
//package(outside all functions, can be used everywhere in the same package)
//global(outside all functions, can be used everywhere accross all packages) -- uppercase letter
