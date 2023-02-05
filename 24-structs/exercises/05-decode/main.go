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
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type exportGames struct {
	Id    int
	Name  string
	Genre string
	Price int
}

type JSONdata struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

const options = `
> list: lists all the games
> id N: queries a game by id
> save: exports the data to json and quits
> quit: quits the program
`

func main() {
	// use your solution from the previous exercise
	var (
		decodedJSON []JSONdata
		in          = bufio.NewScanner(os.Stdin)
	)

	fmt.Printf("Vatsal's game store has %d decodedJSON.\n", len(decodedJSON))
	fmt.Println(options)

	err := json.Unmarshal([]byte(data), &decodedJSON)
	if err != nil {
		fmt.Println("ERR:", err)
	}

	for in.Scan() {
		query := strings.ToLower(in.Text())

		switch {
		case query == "list":
			fmt.Println()
			for i, v := range decodedJSON {
				fmt.Printf("#%d: %-15q (%s) $%-10d\n", i+1, v.Name, v.Genre, v.Price)
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
			} else if id > len(decodedJSON) {
				fmt.Println("\nSorry, we don't have that game!")
				fmt.Println(options)
				break
			}
			fmt.Printf("\n#%d: %-15q (%s) $%-10d\n", id, decodedJSON[id-1].Name, decodedJSON[id-1].Genre, decodedJSON[id-1].Price)
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
