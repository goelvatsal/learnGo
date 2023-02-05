// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
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
	gow := game{item{1, "god of war", 50}, "action adventure"}
	xc := game{item{2, "x-com 2", 30}, "strategy"}
	minecraft := game{item{3, "minecraft", 20}, "sandbox"}
	games := []game{gow, xc, minecraft}

	fmt.Printf("Vatsal's game store has %d games.\n\n", len(games))

	for i, v := range games {
		fmt.Printf("#%d: %-15q (%s) $%-10d\n", i+1, v.name, v.genre, v.price)
	}
}
