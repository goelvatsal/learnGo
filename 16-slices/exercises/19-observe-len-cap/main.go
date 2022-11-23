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
	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//var games []string
	//
	// 2. print the length and capacity of the games slice
	//fmt.Println(len(games))
	//fmt.Println(cap(games))
	//
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//games := []string{}
	////
	//// 4. print the length and capacity of the games slice
	//fmt.Println("len:", len(games), "cap:", cap(games))
	////
	//// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//games = append(games, "pacman", "mario", "tetris", "doom")
	////
	//// 6. print the length and capacity of the games slice
	//fmt.Println("len:", len(games), "cap:", cap(games))
	//
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
	games := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	for i := range games {
		l := games[:i]
		fmt.Println("len:", len(l), "cap:", cap(l))
	}
	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	// for ... {
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", ...)
	// }

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]
	// 2. print the games and the new slice's len and cap
	fmt.Println("len:", len(zero), "cap:", cap(zero))
	//
	// 3. append a new element to the new slice
	zero = append(zero, "hello")
	//
	// 4. print the new slice's lens and caps
	fmt.Println("len:", len(zero), "cap:", cap(zero))
	//
	// 5. repeat the last two steps 5 times (use a loop)
	for i := 0; i < 5; i++ {
		zero = append(zero, "hello")
		fmt.Println("len:", len(zero), "cap:", cap(zero))
	}
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	zero = append(zero, []string{"ultima", "dagger", "pong", "coldspot", "zetra"}...)
	fmt.Println()

	// zero := ...
	// fmt.Printf("games's     len: %d cap: %d\n", ...)
	// fmt.Printf("zero's      len: %d cap: %d\n", ...)

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	for i := range zero {
		v := zero[:i]
		fmt.Println("len:", len(v), "cap:", cap(v))
	}
	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	zero = games[:0]
	for i := 0; i < cap(zero); i++ {
		zero = append(zero, "hello")
		fmt.Println("len:", len(zero), "cap:", cap(zero))
	}
	fmt.Println()

	// zero = ...
	// for ... {
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", ...)
	// }

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[2] = "hi"
	//
	// 2. change the same element of the games slice
	games[2] = "hi"
	//
	// 3. print the games and the zero slices
	s.PrintBacking = true
	s.Show("zero slice", zero)
	s.Show("games slice", games)
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	sl := games[0:9]
	fmt.Println(sl)
}
