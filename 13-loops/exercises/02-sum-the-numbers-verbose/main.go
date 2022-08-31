package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

const max = 10

func main() {
	var sum int

	for i := 1; i <= max; i++ {
		sum += i

		fmt.Printf("%d ", i)
		if i != max {
			fmt.Printf("+ ")
		}
	}
	fmt.Println("=", sum)
}
