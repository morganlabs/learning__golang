package main

import (
	"fmt"
	"sort"
)

func main() {
	// INTS
	ages := []int{16, 23, 45, 76, 29, 40}

	// The sorts.Ints method sorts integers from smallest to largest
	// keep in mind that this .Ints method DOES
	// alter the original slice/array.
	sort.Ints(ages)
	fmt.Println(ages)

	// We can also use .SearchInts to find the index of an integer
	index45 := sort.SearchInts(ages, 45)
	fmt.Println(index45)

	// However, if we try to search for an int that doesn't exist...
	index100 := sort.SearchInts(ages, 100)
	fmt.Println(index100)
	// This will print 6, which is one more than the length of the array

	// And if we append to the array and try again...
	ages = append(ages, 32)
	index100 = sort.SearchInts(ages, 100)
	fmt.Println(index100)
	// This will now print 7, which is one more than the length of the array.
	// If the index of an int cannot be found, it will always return
	// len(arr) + 1

	// STRINGS
	names := []string{"Morgan", "Nagrom", "Ganmor", "Romnag"}

	// We can now use the .Strings method to sort this slice into
	// alphabetical order. Again, this modifies the original array.
	sort.Strings(names)
	fmt.Println(names)

	// The method .SearchStrings also exist, and those are
	// pretty self-explanatory
}
