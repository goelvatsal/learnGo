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
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	var (
		in          = bufio.NewScanner(os.Stdin)
		totalUnique int
		total       int
		unique      = map[string]bool{}
	)

	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		_, ok := unique[word]
		if ok {
			totalUnique--
		}

		totalUnique++
		total++
		unique[word] = true
	}
	fmt.Printf("There are a total of %d words, %d of them are unique.\n", total, totalUnique)
}
