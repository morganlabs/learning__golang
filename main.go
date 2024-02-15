package main

import (
	"fmt"
)

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("This person is a legal adult.")
	} else {
		fmt.Println("This person is a minor.")
	}
}
