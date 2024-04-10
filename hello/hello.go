package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// setup logger
	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	names := []string{"Ania", "Paweł", "Dominik", "Ewa"}
	hellos, err := greetings.GreetMultiplePeople(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, greet := range hellos {
		fmt.Println(greet)
	}
}
