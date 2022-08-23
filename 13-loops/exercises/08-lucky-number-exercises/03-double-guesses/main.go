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
// EXERCISE: Double Guesses
//
//  Let the player guess1 two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guess1ed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess1 one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])

	if err1 != nil {
		fmt.Println("Not a number.")
		return
	}

	guess2, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	min := guess1
	if guess1 < guess2 {
		min = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(min + 1)

		if n == guess1 || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
