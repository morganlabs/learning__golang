package main

import "fmt"

func main() {
	fmt.Print("Hello,")          // This prints WITHOUT a newline at the end. This will be printed on the same line as ...
	fmt.Print(" world.\n")       // If we want to end with a newline, we need to add \n to the end. Or use...
	fmt.Println("Hello, world.") // This, however, adds a newline

	// Concatenation
	name := "Morgan"
	age := 16
	fmt.Println("Hello! My name is", name, "and I am", age, "years old.") // Go automatically adds spaces
	// Or we can use Printf (if you want a newline at the end you must add it yourself):
	// %_ is a "format specifier"
	fmt.Printf("Hello! My name is %v and I am %v-years-old.\n", name, age) // %v is a variable placeholder
	fmt.Printf("Hello! My name is %q and I am %v-years-old.\n", name, age) // %q is a placeholder that surrounds the variables with quotation marks*
	fmt.Printf("name is of type %T\n", name)                               // %T gives you the type of variable
	fmt.Printf("age is of type %T\n", age)

	score := 492.55
	fmt.Printf("You scored %f points!\n", score)    // %f is the qualifier for a float, however it will be followed by multiple zeroes
	fmt.Printf("You scored %0.2f points!\n", score) // We can limit the number of zeroes with %0.xf
	fmt.Printf("You scored %0.1f points!\n", score) // Rounding even works!
	// * This will not work properly with anything other than strings.

	// We can also use the Sprintf function.
	// However, instead of printing to the console, the output gets saved to a variable
	// Like template literals, kinda
	firstName := "Morgan"
	lastName := "Jones"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	fmt.Println(fullName)
}
