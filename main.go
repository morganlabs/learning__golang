package main

import "fmt"

func main() {
	var nameOne string = "Morgan" // Explicitly state type
	var nameTwo = "Nagrom"        // Infer type
	var nameThree string          // Uninitialised variable, type MUST be specified
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Oranmg"
	nameThree = "Ganmor"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Romnag" // A shorthand for initialising a variable w/o declaring it's type
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
}
