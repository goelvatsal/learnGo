package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	if len(os.Args)-1 != 2 {
		fmt.Println("Please pass min and max integers.")
		return
	}

	var sum int

	min, errMin := strconv.Atoi(os.Args[1])
	max, errMax := strconv.Atoi(os.Args[2])

	if min > max {
		fmt.Printf("The max integer is less than the min integer.\n")
		return
	}

	if errMin != nil || errMax != nil {
		fmt.Printf("Either %q or %q is not an integer.\n", os.Args[1], os.Args[2])
		return
	}

	for i := min; i <= max; i++ {
		sum += i

		fmt.Printf("%d ", i)
		if i != max {
			fmt.Printf("+ ")
		}
	}
	fmt.Println("=", sum)
}
