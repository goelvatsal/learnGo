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
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted

	if len(os.Args) == 1 {
		fmt.Println("Type a Hogwarts house name.")
		return
	}

	gryffindor := map[string]string{
		"weasley":    "gryffindor",
		"hagrid":     "gryffindor",
		"dumbledore": "gryffindor",
		"lupin":      "gryffindor",
	}

	hufflepuff := map[string]string{
		"wenlock":   "hufflepuff",
		"scamander": "hufflepuff",
		"helga":     "hufflepuff",
		"diggory":   "hufflepuff",
	}

	ravenclaw := map[string]string{
		"flitwick":    "ravenclaw",
		"bagnold":     "ravenclaw",
		"wildsmith":   "ravenclaw",
		"montmorency": "ravenclaw",
	}

	slytherin := map[string]string{
		"horace":   "slytherin",
		"nigellus": "slytherin",
		"higgs":    "slytherin",
		"scorpius": "slytherin",
	}

	printStudents("gryffindor", gryffindor)
	printStudents("hufflepuff", hufflepuff)
	printStudents("ravenclaw", ravenclaw)
	printStudents("slytherin", slytherin)
}

func printStudents(house string, houseMap map[string]string) {
	if os.Args[1] == house {
		keys := make([]string, 0, len(houseMap))
		for key := range houseMap {
			keys = append(keys, key)
		}

		sort.Strings(keys)
		fmt.Printf(os.Args[1] + " students:\n\n")
		for _, v := range keys {
			fmt.Println(v)
		}
	}
}
