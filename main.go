package main

import "fmt"

func main() {
	greeting1 := greet("Morgan", 16)
	greeting2 := greet("Jane", 19)

	fmt.Println(greeting1)
	fmt.Println(greeting2)

	farewell1 := farewell("Morgan")
	farewell2 := farewell("Jane")

	fmt.Println(farewell1)
	fmt.Println(farewell2)
}

func greet(name string, age int) string {
	return fmt.Sprintf("Hello, %v! You're %v-years-old, correct?", name, age)
}

func farewell(name string) string {
	return fmt.Sprintf("Goodbye, %v! Hope to see you soon!", name)
}
