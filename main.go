package main

import "fmt"

func main() {
	// 1.1. Define a map
	menu := map[string]float64{
		"pie":        5.99,
		"fudge cake": 3.99,
	}

	// 1.2. Print the map
	fmt.Println(menu)

	// 1.3. Pass it into a function
	modifyPrice(menu)

	// 1.4. Print it out
	fmt.Println(menu)

	// Unlike in part 1, the value is actually modified this time!
	// This is because certain data structures, such as maps, can be
	// extremely large, and it is more memory efficient and quicker to
	// instead just give the function a pointer to the original value
	// instead of copying the entire value

	// Data structures that would be *copied* would be:
	// * strings
	// * ints
	// * floats
	// * booleans
	// * arrays
	// * structs
	// Data structures that would pass in a *pointer* would be:
	// * slices
	// * maps
	// * functions
}

// 1.3.1. Take in a map variable
func modifyPrice(menu map[string]float64) {
	// 1.3.2. Modify the map
	menu["pie"] = 6.99

	// 1.3.3. Return nothing
}
