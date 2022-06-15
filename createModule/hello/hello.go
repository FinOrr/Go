package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including the log entry
	// prefix and a flag to disable printing the time, source file and
	// line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, get a greeting message and print it to
	// the terminal.
	// message := greetings.Hello("Random Stranger") <= this line commented out for testing the error handling
	fmt.Println(message)
}
