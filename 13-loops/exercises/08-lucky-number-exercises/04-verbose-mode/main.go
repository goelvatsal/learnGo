// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

The NEW -v option allows the program to display the 5 random numbers! Just type -v before your guess!

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var (
		guess   int
		err     error
		verbose bool
	)

	if args[0] == "-v" {
		guess, err = strconv.Atoi(args[1])
		verbose = true
	} else {
		guess, err = strconv.Atoi(args[0])
	}

	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	var list []int
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if verbose {
			if n == guess {
				list = append(list, n)
			} else if n != guess {
				list = append(list, n)
			}
		} else if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	var counter int
	for i := 0; i != len(list); i++ {
		if list[i] == guess {
			counter += 1
		}
	}

	if counter > 0 {
		fmt.Println(list, "🎉  YOU WIN!")
	} else {
		fmt.Println(list, "☠️  YOU LOST... Try again?")
	}
}
