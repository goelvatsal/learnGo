package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year, month := time.Now().Year(), os.Args[1]

	switch m := strings.ToLower(month); m {
	case "april", "june", "september", "november":
		fmt.Printf("%q has 30 days.\n", month)
	case "january", "march", "may", "july", "august", "october", "december":
		fmt.Printf("%q has 31 days.\n", month)
	case "february":
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			fmt.Printf("%q has 29 days.\n", month)
		} else {
			fmt.Printf("%q has 28 days.\n", month)
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
	}
}
