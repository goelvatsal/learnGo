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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
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

const options = `
> list: lists all the games
> id N: queries a game by id
> quit: quits the program
`

func main() {
	// use your solution from the previous exercise
	var (
		gow       = game{item{1, "god of war", 50}, "action adventure"}
		xc        = game{item{2, "x-com 2", 30}, "strategy"}
		minecraft = game{item{3, "minecraft", 20}, "sandbox"}
		games     = []game{gow, xc, minecraft}
		in        = bufio.NewScanner(os.Stdin)
	)

	fmt.Printf("Vatsal's game store has %d games.\n", len(games))
	fmt.Println(options)

	for in.Scan() {
		query := strings.ToLower(in.Text())

		switch {
		case query == "list":
			fmt.Println()
			for i, v := range games {
				fmt.Printf("#%d: %-15q (%s) $%-10d\n", i+1, v.name, v.genre, v.price)
			}
			fmt.Println(options)
		case query == "quit":
			fmt.Println("\nbye!")
			return
		case strings.Contains(query, "id"):
			contents := strings.Fields(query)
			if len(contents) != 2 {
				fmt.Println("\nwrong id")
				fmt.Println(options)
				break
			}

			id, err := strconv.Atoi(contents[1])
			if err != nil {
				fmt.Println("wrong id")
				fmt.Println(options)
				break
			} else if id > len(games) {
				fmt.Println("\nSorry, we don't have that game!")
				fmt.Println(options)
				break
			}
			fmt.Printf("\n#%d: %-15q (%s) $%-10d\n", id, games[id-1].name, games[id-1].genre, games[id-1].price)
			fmt.Println(options)
		default:
			return
		}
	}
}
