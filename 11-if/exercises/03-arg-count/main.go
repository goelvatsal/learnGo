package main

import "fmt"
import "os"

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	lenOs := len(os.Args) - 1

	if lenOs == 0 {
		fmt.Println("Give me args")
	} else if lenOs == 1 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if lenOs == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", os.Args[1], os.Args[2])
	} else if lenOs > 2 {
		fmt.Printf("There are %d arguments.\n", lenOs)
	}
}
