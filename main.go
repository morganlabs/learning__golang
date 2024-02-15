package main

import (
	"fmt"
)

func main() {
	a := 0
	for a <= 5 { // The same as while (a <= 5) {}
		fmt.Printf("a = %v\n", a)
		a++
	}

	for b := 6; b <= 10; b++ { // The same as for (let b = 6; b <= 10; b++) {}
		fmt.Printf("b = %v\n", b)
	}

	// Iterating through arrays & slices
	names := []string{"Morgan", "Nagrom", "Ganrom", "Romnag"}

	// Method 1
	for i := 0; i < len(names); i++ {
		name := names[i]
		fmt.Printf("names[%v] = %v\n", i, name)
	}

	// Or we can iterate through an array/slice with...
	// Method 2
	for i, name := range names {
		fmt.Printf("names[%v] = %v\n", i, name)
	}

	// And if we don't need index OR the value, replace it with an underscore
	for _, name := range names {
		fmt.Println(name)

		// NOTE: If we change the value of the array item from within a loop,
		// it does NOT modify the original array.
		name = "Someone Else"
	}

	// This is unaffected
	fmt.Println(names)
}
