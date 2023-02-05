// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import (
	"fmt"
)

func main() {
	a, b := 6.28, 3.14
	fmt.Printf("before: a: %g ----- b: %g\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("after: a: %g ----- b: %g\n", a, b)

	fa, fb := &a, &b
	fmt.Printf("\nfa: %p (%g) ----- fb: %p (%g)\n", fa, *fa, fb, *fb)
	ra, rb := swapAddress(fa, fb)
	fmt.Printf("\nfa: %p (%g) ----- fb: %p (%g)\n", ra, *ra, rb, *rb)

}

func swap(a, b float64) (float64, float64) {
	la, lb := &a, &b
	la, lb = lb, la

	return *la, *lb
}

func swapAddress(a, b *float64) (*float64, *float64) {
	la, lb := &a, &b
	la, lb = lb, la

	return *la, *lb
}
