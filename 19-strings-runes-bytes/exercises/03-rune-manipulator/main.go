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
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = words

	for _, word := range words {
		c := utf8.RuneCountInString(word)
		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("%q\n", word)
		fmt.Printf("\thas %d bytes and %d runes\n", len(word), c)

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes: % x\n", word)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\trunes: ")
		for _, r := range word {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Printf("\trunes: ")
		for _, r := range word {
			fmt.Printf("%q ", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(word)
		fmt.Printf("\tfirst: %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		last, lastSize := utf8.DecodeLastRuneInString(word[c-1:])
		fmt.Printf("\tlast: %q (%d bytes)\n", last, lastSize)

		// Slice and print the first two runes of the strings
		runeSlice := []rune(word)
		fmt.Printf("\tfirst 2: %q\n", runeSlice[:2])

		// Slice and print the last two runes of the strings
		fmt.Printf("\tlast 2: %q\n", runeSlice[c-2:])

		// Convert the string to []rune
		// Print the first and last two runes
		fmt.Printf("\tfirst 2: %q\n", runeSlice[:2])
		fmt.Printf("\tlast 2: %q\n", runeSlice[c-2:])
	}
}
