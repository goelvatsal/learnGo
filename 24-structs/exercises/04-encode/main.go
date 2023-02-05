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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > Id N   : queries a game by Id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "Id": 1,
//                  "Name": "god of war",
//                  "Genre": "action adventure",
//                  "Price": 50
//          },
//          {
//                  "Id": 2,
//                  "Name": "x-com 2",
//                  "Genre": "strategy",
//                  "Price": 40
//          },
//          {
//                  "Id": 3,
//                  "Name": "minecraft",
//                  "Genre": "sandbox",
//                  "Price": 20
//          }
//  ]
//
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

type exportGames struct {
	Id    int
	Name  string
	Genre string
	Price int
}

const options = `
> list: lists all the games
> Id N: queries a game by Id
> save: exports the data to json and quits
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
		case strings.Contains(query, "Id"):
			contents := strings.Fields(query)
			if len(contents) != 2 {
				fmt.Println("\nwrong Id")
				fmt.Println(options)
				break
			}

			id, err := strconv.Atoi(contents[1])
			if err != nil {
				fmt.Println("wrong Id")
				fmt.Println(options)
				break
			} else if id > len(games) {
				fmt.Println("\nSorry, we don't have that game!")
				fmt.Println(options)
				break
			}
			fmt.Printf("\n#%d: %-15q (%s) $%-10d\n", id, games[id-1].name, games[id-1].genre, games[id-1].price)
			fmt.Println(options)
		case query == "save":
			exportGamesData := []exportGames{
				{1, "god of war", "action adventure", 50},
				{2, "x-com 2", "strategy", 40},
				{3, "minecraft", "sandbox", 20},
			}

			out, err := json.MarshalIndent(exportGamesData, "", "\t")
			if err != nil {
				fmt.Println("ERR:", err)
			}
			fmt.Println(string(out))
			return
		default:
			return
		}
	}
}
