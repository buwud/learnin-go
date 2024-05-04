package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"buwu",
		"evrenkasifi",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	} //else println
	fmt.Println(messages)
}
