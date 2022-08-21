package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user1    = "jack"
	pass1    = "1888"
	user2    = "inanc"
	pass2    = "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user1 && u != user2 {
		fmt.Printf(errUser, u)
	} else if p != pass1 && p != pass2 {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
