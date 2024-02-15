package main

import "fmt"

func main() {
	names := []string{"Morgan", "Jane", "John"}

	cycleNames(names, greet)
	cycleNames(names, farewell)
}

// A function can take in another function as an argument (a "callback function")
// This callback MUST match the same arguments in as the function being called upon
// So this means that cycleNames takes in a slice of names, and a function that MUSt
// ONLY take in 1 string as an argument. Then, this callback can be called from
// within the function!
// Callback function canNOT return any values.
func cycleNames(names []string, callback func(string)) {
	for _, name := range names {
		callback(name)
	}
}

func greet(name string) {
	fmt.Printf("Hello, %v! How are you?\n", name)
}

func farewell(name string) {
	fmt.Printf("Goodbye, %v! Hope to see you soon!\n", name)
}
