package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Morgan"
	greeting := fmt.Sprintf("Hello, %v!", name)

	// NOTE: All of these methods are NOT destructive, as in they
	// do NOT modify the original string input. Thee values must
	// either be saved to a new variable, printed or saved over
	// the original variable to be useful

	// We can import "strings" and use strings.Contains to find out
	// if a string contains a substring.
	// Here we are checking if `greeting` contains `name`, which will
	// always be true.
	containsName1 := strings.Contains(greeting, name)

	// But here, we're checking if "Jack" exists. This returns false
	containsName2 := strings.Contains(greeting, "Jack")

	// However, if we change the name to Jack and reformat the greeting
	// string...
	name = "Jack"
	greeting = fmt.Sprintf("Hello, %v!", name)

	// This returns true
	containsName3 := strings.Contains(greeting, "Jack")
	fmt.Println(containsName1, containsName2, containsName3)

	// We can also use strings.Replace and .ReplaceAll to replace either
	// n or all instances of a substring inside of a string
	// .Replace requires you to specify the number of instances to replace
	// at the end of the method
	noMoreJack := strings.Replace(greeting, "Jack", "Morgan", 1)
	fmt.Println(noMoreJack)

	// And we can use .ReplaceAll to replace every instance
	greeting = "Hello, Earth. Hello, Earth. Hello, Earth."
	noMoreEarth := strings.ReplaceAll(greeting, "Earth", "Mars")
	fmt.Println(noMoreEarth)

	// There's also .ToUpper and .ToLower to convert a string to
	// upper- or lower-case respectively
	noMoreEarthShouting := strings.ToUpper(noMoreEarth)
	noMoreEarthWhisper := strings.ToLower(noMoreEarth)
	fmt.Println(noMoreEarthShouting)
	fmt.Println(noMoreEarthWhisper)

	// Or .Index, which allows us to find the index of a substring
	whereIsMars := strings.Index(noMoreEarth, "Mars")
	fmt.Println(whereIsMars)
	// This will print 7, as the first instance of Mars is found at index 7.

	// Or we can use .Split to split a string into a slice
	sliceMars := strings.Split(noMoreEarth, " ")
	fmt.Println(sliceMars)
}
