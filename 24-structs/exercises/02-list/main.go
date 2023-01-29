// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	// use your solution from the previous exercise
	var (
		gow       = game{item{1, "god of war", 50}, "action adventure"}
		xc        = game{item{2, "x-com 2", 30}, "strategy"}
		minecraft = game{item{3, "minecraft", 20}, "sandbox"}
		games     = []game{gow, xc, minecraft}
		in        = bufio.NewScanner(os.Stdin)
	)

	fmt.Printf("Vatsal's game store has %d games.\n\n", len(games))

	fmt.Println("> list: lists all the games")
	fmt.Println("> quit: quits the program", "\n")

	for in.Scan() {
		query := strings.ToLower(in.Text())

		switch query {
		case "list":
			fmt.Println()
			for i, v := range games {
				fmt.Printf("#%d: %-15q (%s) $%-10d\n", i+1, v.name, v.genre, v.price)
			}

			fmt.Println("\n> list: lists all the games")
			fmt.Println("> quit: quits the program")
		case "quit":
			fmt.Println("\nbye!")
			return
		default:
			return
		}
	}
}
