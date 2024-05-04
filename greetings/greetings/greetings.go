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

// randomly return a string from slice
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello beautiful %v!",
		"Hail, %v. Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
