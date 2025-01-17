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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	if len(os.Args)-1 >= 6 || len(os.Args)-1 == 0 {
		fmt.Println("Please enter a maximum of 5 numbers.")
		return
	}

	var (
		arr     [5]int
		avg     int
		lenArgs int
	)

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err == nil {
			lenArgs += 1
			avg += n
		}
		arr[i-1] = n
	}
	avg = avg / lenArgs
	fmt.Printf("Your numbers: %v\n", arr)
	fmt.Println("Average:", avg)
}
