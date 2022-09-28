package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	names[0] = "Friend1"
	names[1] = "Friend2"
	names[2] = "Friend3"
	distances[0] = 2
	distances[1] = 10
	distances[2] = 5
	distances[3] = 20
	distances[4] = 25
	data[0] = 72
	data[1] = 19
	data[2] = 63
	data[3] = 10
	data[4] = 47
	ratios[0] = 1.57
	alives[0] = true
	alives[1] = true
	alives[2] = false
	alives[3] = false

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
