package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [command] [string]")
		return
	}

	c, s := os.Args[1], os.Args[2]

	switch c {
	case "title":
		fmt.Printf("Output: %q\n", strings.Title(s))
	case "upper":
		fmt.Printf("Output: %q\n", strings.ToUpper(s))
	case "lower":
		fmt.Printf("Output: %q\n", strings.ToLower(s))
	default:
		fmt.Printf("Unknown command: %q\n", c)
		return
	}
}
