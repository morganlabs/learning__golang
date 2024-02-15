package main

import "fmt"

func main() {
	// Maps
	// These are similar to Python dictionaries or, to a lesser
	// extent, JavaScript objects.
	// The type definition of a map goes as follows:
	// If you want the keys to be of type string, and the values of type
	// float64, then do the following:
	menu := map[string]float64{
		"soup":        5.99,
		"chicken pie": 7.99,
		"salad":       6.99,
		"fudge cake":  3.99, // Trailing commas are REQUIRED
	}
	// So the typedef goes map[typeof keys]typeof values

	fmt.Println(menu)
	fmt.Println(menu["soup"]) // Bracker syntax is needed to select a value.

	// Loop
	for item, price := range menu {
		fmt.Printf("The %v costs £%v\n", item, price)
	}

	// Ints as key type
	phonebook := map[int]string{
		12345678901: "Morgan Jones",
		23456789012: "John Doe",
		34567890123: "Jane Doe",
		45678901234: "Jack Lastname",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[12345678901])
	phonebook[12345678901] = "Senoj Nagrom"
	fmt.Println(phonebook[12345678901])
}
