package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received by the function, then return
	// a value that embeds the name in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	// Uncomment to force an error, that will be detected in testing:
	//message := fmt.Sprint(randomFormat())
	return message, nil // Don't forget to return nil error
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
// >> map[key-type]value-type
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// Inn the map, associate the retreived message with the name
		messages[name] = message
	}
	return messages, nil
}

// init set the initial values for the variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages
// The returned message is selected at random
func randomFormat() string {
	// A slice of message formats; like a dynamically resized array
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Now return a randomly selected message format by
	// specifying a random index for the slice of formats
	return formats[rand.Intn((len(formats)))]
}
