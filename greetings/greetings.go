package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func GreetMultiplePeople(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomGreeting() string {
	// define a list of greetings
	greetings := []string{
		"Hi %v. Welcome!\n",
		"Great to see you %v\n",
		"Dzień dobry %v\n",
		"Guten tag %v\n",
	}
	return greetings[rand.Intn(len(greetings))]
}
