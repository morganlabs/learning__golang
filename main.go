package main

import "fmt"

func main() {
	// Initialise an array
	// Arrays are **fixed length**. You cannot add more to an array
	// than the length described at initialisation, however you may
	// have less items inside. And, of course, you cannot deviate
	// from the type you described. You cannot have an array of
	// strings and ints

	// Explicitly type an array
	var ages1 [6]int = [6]int{16, 23, 38, 42, 59, 60}

	// Infer the type (and length) of an array
	var ages2 = [6]int{16, 23, 38, 42, 59, 60}
	ages3 := [6]int{16, 23, 38, 42, 59, 60}

	// Initialise an empty array, ready for data
	// If this is printed without any data, the number 0 will be
	// repeated for the length of the array.
	var ages4 [6]int = [6]int{}

	// Chage the values inside an array
	// This works
	ages2[0] = 17

	// However, this does not, as the array only goes up to index 6
	// ages2[7] = 17

	fmt.Println(ages1, ages2, ages3, ages4)

	// Print the length of an array using the len() method
	fmt.Println(len(ages1))

	// Slides
	// These use the array type under the hood.
	// Slices are of a dynamic length, unlike an array
	var scores = []int{100, 200, 300, 400}

	// We can modify the values inside of a slice the same way we do with an array
	scores[1] = 250

	// However, because they are also of dynamic lengths, we can append to a slice, too
	// Keep in mind that this append function does NOT modify the original slice,
	// but instead return a new slice to us. This means we need to assign the new
	// slice to a new variable, or the original variable if we want to modify it
	scores = append(scores, 500)
	newScores := append(scores, 600)

	fmt.Println(scores, newScores)

	// Ranges
	// These let you grab parts of an array or slice from index X to index Y.
	// Remember that index X is INclusive and index Y is EXclusive
	// This grabs ages1[1] and ages[2], but not ages[3] as index Y is EXclusive
	rangeOne := ages1[1:3]

	// This grabs everything STARTING from index 2 all the way until the end of the array
	rangeTwo := ages1[2:]

	// This grabs everything from the beginning of the array until index 4 (remember, the second number is EXclusive)
	rangeThree := ages1[:5]

	// And because ranges are Slices by default (reguardless of whether they were derived from
	// an array or a slice), we can append to them
	rangeThree = append(rangeThree, 48)

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
