package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
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
		fmt.Printf("%q is not an integer.\n", args[0])
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not an integer.\n", args[1])
		return
	}

	if max < min {
		fmt.Printf("The max integer is less than the min integer.\n")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		isEven := i%2 == 0
		if i != max && isEven {
			fmt.Printf("%d ", i)

			if i != max-1 {
				fmt.Printf("+ ")
			}
		} else if !isEven {
			continue
		}

		sum += i
	}

	// what is being done here?
	if max%2 == 0 {
		fmt.Printf("%d ", max)
	}

	fmt.Println("=", sum)
}
