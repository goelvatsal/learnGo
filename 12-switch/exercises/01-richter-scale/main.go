package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	n, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Printf("%s is not a number.\n", os.Args[1])
		return
	}

	switch {
	case n < 2.0:
		fmt.Printf("%g is a micro earthquake.\n", n)
	case n < 3.0 && n >= 2.0:
		fmt.Printf("%g is a very minor earthquake.\n", n)
	case n < 4.0 && n >= 3.0:
		fmt.Printf("%g is a minor earthquake.\n", n)
	case n < 5.0 && n >= 4.0:
		fmt.Printf("%g is a light earthquake.\n", n)
	case n < 6.0 && n >= 5.0:
		fmt.Printf("%g is a moderate earthquake.\n", n)
	case n < 7.0 && n >= 6.0:
		fmt.Printf("%g is a strong earthquake.\n", n)
	case n < 8.0 && n >= 7.0:
		fmt.Printf("%g is a major earthquake.\n", n)
	case n < 10.0 && n >= 7.0:
		fmt.Printf("%g is a great earthquake.\n", n)
	case n >= 10:
		fmt.Printf("%g is a massive earthquake.\n", n)
	}
}
