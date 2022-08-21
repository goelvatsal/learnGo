package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	firstname := "Vatsal"
	lastname := "Goel"

	fmt.Printf("My name is %s and my last name is %s\n", firstname, lastname)
}
