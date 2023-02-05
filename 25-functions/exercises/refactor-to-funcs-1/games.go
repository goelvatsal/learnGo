package main

import (
	"fmt"
	"sort"
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

func newGame(id, price int, name, genre string) game {
	tempItem := item{id, name, price}
	return game{tempItem, genre}
}

func addGame(gameSlice []game, gameNorm game) []game {
	temp := append([]game{}, []game{gameNorm}...)
	temp = append(temp, gameSlice...)

	return temp
}

func load() (gameSlice []game) {
	gameSlice = addGame(gameSlice, newGame(1, 50, "god of war", "action adventure"))
	gameSlice = addGame(gameSlice, newGame(2, 40, "x-com 2", "strategy"))
	gameSlice = addGame(gameSlice, newGame(3, 20, "minecraft", "sandbox"))
	sort.Slice(gameSlice, func(i, j int) bool { return gameSlice[i].id < gameSlice[j].price })

	return gameSlice
}

func indexByID(gameSlice []game) (gameMap map[int]game) {
	gameMap = map[int]game{}
	for _, g := range gameSlice {
		gameMap[g.id] = g
	}
	return gameMap
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
}
