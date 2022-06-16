package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry
	// prefix and a flag to disable printing the time, source file and
	// line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names, where each name is a key-value
	names := []string{
		"Gary",
		"Samantha",
		"Putek",
	}

	// Request a greeting message

	// To return just a single greeting for Gary:
	// message, err := greetings.Hello("Gary")
	// To return a greeting for each of the names passed in:
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, get a greeting message and print it to
	// the terminal.
	// message := greetings.Hello("Random Stranger") <= this line commented out for testing the error handling
	fmt.Println(messages)
}
