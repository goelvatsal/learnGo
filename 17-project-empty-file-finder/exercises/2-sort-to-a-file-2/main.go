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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     1. apple
//     2. banana
//     3. orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + You can use strconv.AppendInt function to append an int
//     to a byte slice. strconv contains a lot of functions for appending
//     other basic types to []byte slices as well.
//
//   + You can append individual characters to a byte slice using
//     rune literals (because: rune literal are typeless numerics):
//
//     var slice []byte
//     slice = append(slice, 'h', 'i', ' ', '!')
//     fmt.Printf("%s\n", slice)
//
//     Above code prints: hi !
// ---------------------------------------------------------

func main() {
	l := len(os.Args) - 1
	if l == 0 {
		fmt.Println("Give me values.")
		return
	}

	args := os.Args[1:]
	sort.Strings(args)
	var out []byte
	counter := 1

	for _, v := range args {
		out = strconv.AppendInt(out, int64(counter), 10)
		out = append(out, '.', ' ')
		out = append(out, v...)
		out = append(out, "\n"...)
		counter++
	}

	err := os.WriteFile("results.txt", out, 0755)
	if err != nil {
		fmt.Println(err)
	}
}
