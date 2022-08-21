package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Please pass min and max integers.")
		return
	}

	min, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Either %q or %q is not an integer.\n", args[0], args[1])
	}

	max, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("Either %q or %q is not an integer.\n", args[0], args[1])
	}

	if min > max {
		fmt.Printf("The max integer is less than the min integer.\n")
		return
	}

	i := min
	var sum int
	for {
		if i >= max {
			break
		}

		isEven := i%2 == 0

		if i != max && isEven {
			fmt.Printf("%d ", i)

			if i != max-1 {
				fmt.Printf("+ ")
			}
		}

		if isEven {
			sum += i
		}
		i++
	}

	if max%2 == 0 {
		fmt.Printf("%d ", max)
		sum += i
	}

	fmt.Println("=", sum)
}
