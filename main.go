// The `package main` line indicates we are creating a package named main.
// This acts as the entry package to our file.
package main

import "fmt"

// There can only be ONE functon called `main` in an entire go program.
// This is the beginning of our program. `main` cannot exist in any other
// files.
func main() {
	// Here, we're going to use the function `greet` which is defined in greetings.go
	// which is under the same package name `main`.
	greet("Morgan")

	// And we're also going to access the `points` slide, which is also defined in
	// greetings.go
	for _, points := range points {
		fmt.Println(points)
	}
}

// Currently at this stage in the course I am watching, we're running these go files with
// `go run main.go`
// However, doing this will result in go telling us that `points` and `greet` are not defined.
// To fix this, we need to include every file inside of the `main` package for their vairables
// and functions to be exposed to us:
// `go run main.go greetings.go`
