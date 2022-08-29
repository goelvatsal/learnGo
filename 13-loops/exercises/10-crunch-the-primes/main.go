package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	for _, q := range args {
		c, err := strconv.Atoi(q)
		if err != nil {
			continue
		}

		isPrime(c)
	}
	fmt.Println()
}

func isPrime(n int) int {
	if n == 1 {
	} else if n == 2 {
		fmt.Printf("%d ", n)
	} else if n == 3 {
		fmt.Printf("%d ", n)
	}

	if n%2 == 0 {
		return 0
	} else if n%3 == 0 {
		return 0
	}

	i := 5
	w := 2

	for {
		if i*i <= n {
			if n%i == 0 {
				return 0
			}
		} else {
			break
		}
	}

	i += w
	w = 6 - w
	fmt.Printf("%d ", n)

	return 0
}
