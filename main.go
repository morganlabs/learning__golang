package main

import (
	"fmt"
	"strings"
)

func main() {
	// When using multiple return types, similarly to a loop through an
	// array/slice, you must account for every returned value. Use an
	// underscore to discard the value.
	firstInitial1, lastInitial1 := getInitials("Morgan Jones")
	firstInitial2, _ := getInitials("Morgan")

	fmt.Printf("%v%v\n", firstInitial1, lastInitial1)
	fmt.Printf("%v\n", firstInitial2)
}

// Multiple return values are signified by making the return type:
// (type, type, type, ...)
func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	names := strings.Split(name, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1])
	}

	if len(initials) == 1 {
		return initials[0], "_"
	}

	return initials[0], initials[1]
}
