package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	var (
		names     = [...]string{"Friend1", "Friend2", "Friend3"}
		distances = [...]int{2, 10, 5, 20, 25}
		data      = [...]uint8{72, 19, 63, 10, 47}
		ratios    = [...]float64{1.57}
		alives    = [...]bool{true, true, false, false}
		zero      = [...]uint8{}
	)

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println()

	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()

	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %v\n", i, data[i])
	}
	fmt.Println()

	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %g\n", i, ratios[i])
	}
	fmt.Println()

	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()

	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	// FOR RANGE STARTS HERE
	fmt.Println("FOR RANGE STARTS HERE")
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}
	fmt.Println()

	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}
	fmt.Println()

	for i, v := range data {
		fmt.Printf("data[%d]: %v\n", i, v)
	}
	fmt.Println()

	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %g\n", i, v)
	}
	fmt.Println()

	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}
	fmt.Println()

	for i, v := range zero {
		fmt.Printf("zero[%d]: %q\n", i, v)
	}
}
