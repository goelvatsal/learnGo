// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"reflect"
)

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var null *computer
	// compare the pointer variable to a nil value, and say it's nil
	fmt.Println("is null equal to nil?", null == nil, "\n")
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}
	// put the apple into a new pointer variable
	newApple := apple
	// compare the apples: if they are equal say so and print their addresses
	if newApple == apple {
		fmt.Println("is newApple equal to apple?", newApple == apple)
		fmt.Printf("newApple addr: %v\napple addr: %v\n\n", &newApple, &apple)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}
	sonyAddr := &sony
	// compare apple to sony, if they are not equal say so and print their
	// addresses
	fmt.Println("is apple equal to sony?", reflect.DeepEqual(apple, sony))
	fmt.Println("sony addr:", sonyAddr)
	fmt.Println("apple addr", &apple, "\n")
	// put apple's value into a new ordinary variable
	appleOrd := *apple
	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple's pointer address: %v %v\n", &apple, apple)
	fmt.Printf("appleOrd's pointer address: %v\n\n", &appleOrd)
	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == appleOrd {
		fmt.Println("apple and appleOrd are the same")
	}
	fmt.Println("value that is pointed by the apple:", &apple)
	fmt.Println("value that is pointed by the appleOrd:", &appleOrd, "\n")
	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("apple: %+v ------ appleVal: %+v\n", *apple, appleOrd)
	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Println("sony: ", sony)
	// write a new constructor func that returns a pointer to a computer
	// print the returned value
	fmt.Printf("apple: %+v\n\n", returnResult(apple))

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address: %x\n", newComp("dell"))
	fmt.Printf("lenovo's address: %x\n", newComp("lenovo"))
	fmt.Printf("hp's address: %x\n", newComp("hp"))
}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c *computer, brand string) {
	c.brand = brand
}

// write a func that returns the value that is pointed by the given *computer
func returnResult(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComp(brand string) *computer {
	return &computer{brand: brand}
}
