package main

import "os"
import "fmt"
import "strconv"

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	lenOs, args := len(os.Args)-1, os.Args

	if lenOs < 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	un, pw := "jack", 1888

	if args[1] != un {
		fmt.Printf("Access denied to %q.\n", args[1])
	} else if args[2] != strconv.Itoa(pw) {
		fmt.Printf("Invalid password for %q.\n", args[1])
	} else {
		fmt.Printf("Access granted to %q.\n", args[1])
	}
}
