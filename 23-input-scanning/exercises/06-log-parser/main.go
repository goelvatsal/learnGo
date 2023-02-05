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
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	var (
		in       = bufio.NewScanner(os.Stdin)
		parseMap = map[string]int{}
		total    int
		lineNum  int
	)

	for in.Scan() {
		line := strings.Fields(in.Text())
		domain := line[0]
		if len(line) != 2 {
			fmt.Printf("wrong input: [%s] (line #%d)\n", domain, lineNum+1)
			return
		}

		visits, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("wrong input: \"%s\" (line #%d)\n", line[1], lineNum+1)
			return
		} else if visits < 0 {
			fmt.Printf("wrong input: \"%v\" (line #%d)\n", visits, lineNum+1)
			return
		}

		parseMap[domain] += visits
		total += visits
		lineNum++
	}

	fmt.Printf("%-30s %-30s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 40))

	for k, m := range parseMap {
		fmt.Printf("%-30s %-30d\n", k, m)
	}
	fmt.Printf("\nTOTAL %27d\n", total)
}
