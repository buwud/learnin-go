package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
	//nil means no error
}

// messages according to the names of the people specifically using map
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	//blank identifier -> For instance, when calling a function that returns a value(for me name) and an error, but only the error is important, use the blank identifier to discard the irrelevant value.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil { //if there is an error
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomly return a string from slice
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello beautiful %v!",
		"Hail, %v. Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
