package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	/* if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("So long")
	} else if word == "greetings" {
		fmt.Println("Salutation")
	} else {
		fmt.Println("I do not know what you mean")
	} */

	/* greet := "greetings"
	switch word {
	case "hi":
		fmt.Println("Very informal")
		fallthrough //matches two cases
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "bye":
		fmt.Println("So long")
	case "farewell": // no op case
	case greet: // variables can be used as case values
		fmt.Println("Salutation")
	default:
		fmt.Println("I do not know what you mean")
	} */

	/* greet := "greetings"
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very informal")
		fallthrough //matches two cases
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "bye":
		fmt.Println("So long")
	case "farewell": // no op case
	case greet: // variables can be used as case values
		fmt.Println("Salutation")
	default:
		fmt.Println("I do not know what you mean, but it was", l, "characters long")
	} */

	greet := "greetings"
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very informal")
		fallthrough //matches two cases
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "bye":
		fmt.Println("So long")
	case "farewell": // no op case
	case greet: // variables can be used as case values
		fmt.Println("Salutation")
	default:
		fmt.Println("I do not know what you mean, but it was", l, "characters long")
	}
}
