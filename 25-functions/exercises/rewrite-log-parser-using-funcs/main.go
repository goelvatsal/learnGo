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

func main() {
	p := parser{sum: make(map[string]result)}

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		fieldParse(fields, p)

		domain := fields[0]

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		// Collect the unique domains
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// Keep track of total and per domain visits
		p.total += visits

		// You cannot assign to composite values
		// p.sum[domain].visits += visits

		// create and assign a new copy of `visit`
		newVisitCopy(p, domain, visits)
	}

	// Print the visits per domain
	printTotal(p)

	// Print the total visits for all domains
	printTotalVisits(p)

	// Let's handle the error
	handleError(in)
}
