package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	var (
		names     = [3]string{"Friend1", "Friend2", "Friend3"}
		distances = [5]int{2, 10, 5, 20, 25}
		data      = [5]uint8{72, 19, 63, 10, 47}
		ratios    = [1]float64{1.57}
		alives    = [4]bool{true, true, false, false}
		zero      = [0]uint8{}
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
