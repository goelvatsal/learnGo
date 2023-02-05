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
	"strconv"
	"strings"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))
	return
}

func addGame(games []game, g game) []game {
	return append(games, g)
}

func newGame(id, price int, name, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

func indexByID(games []game) (byID map[int]game) {
	byID = make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}

//1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []game, map[int]game
//        Output: bool
//
//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
//

func runCmd(input string, games []game, byID map[int]game) bool {
	games = load()
	byID = indexByID(games)
	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return false
	}

	switch cmd[0] {
	case "quit":
		cmdQuit()
	case "list":
		cmdList(games)
	case "id":
		cmdByID(cmd, games, byID)
	}
	return true
}

//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

// 3- Add a func that handles the list command:
//
//	Name  : cmdList
//	Inputs: []game
//	Output: bool
func cmdList(gameSlice []game) bool {
	for _, g := range gameSlice {
		printGame(g)
	}
	return true
}

//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool

func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return false
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return false
	}

	games = load()
	byID = indexByID(games)
	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return false
	}

	printGame(g)
	return true
}
