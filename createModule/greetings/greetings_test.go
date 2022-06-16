package greetings

import (
	"regexp"
	"testing"
)

// Running the [ go test -v ] command in the greetings directory will execute
// the test, with the verbose flag (giving more feedback on the tests)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloempty calls greeting.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
