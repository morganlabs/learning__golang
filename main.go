package main

import "fmt"

func main() {
	// Pass-by-value: Part 1
	// Go is a PBV language. This means that whenever we pass a value into a
	// function, it creates a copy of that value.

	// 1.1. Define a variable
	name := "Morgan"

	// 1.2. Print it out
	fmt.Println(name)

	// 1.3. Pass it into a function
	changeName(name)

	// 1.4. The output of the original variable is unchanged
	fmt.Println(name)

	// This is because we Go automatically created a **copy** of the
	// `name` variable for the function to use.

	// But if we WANT to modify the value, we can return a value and
	// assign it to the original variable.
	// 1.1. Define a variable
	name = "Morgan"

	// 2.2. Run the revised function and assign it to the variable
	name = changeNameReturn(name)

	// 2.2. Print the name and notice it has been modified
	fmt.Println(name)
}

// 1.3.1. Take in a **copy** of a name
func changeName(name string) {
	// 1.3.2. Change the **copy** of the name
	name = "Another Name"

	// 1.3.3. Return nothing
}

// 2.1.1. Take in a **copy** of a name with the return type string
func changeNameReturn(name string) string {
	// 2.1.2. Change the **copy** of the name
	name = "Another Name"

	// 2.1.3. Return the copy of the name
	return name
}
