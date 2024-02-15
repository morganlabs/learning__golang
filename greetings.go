// Any other go file that starts with `package main` will be considered
// a part of the main package scope. By doing this, we can share variables
// and functions between all files in `main`.
package main

import "fmt"

// Any variables, functions, etc., that are defined at the **top level**
// of a go package are available in any other file in the same package.
// This does NOT include variables defined within a function, of course.
var points = []int{10, 23, 45, 29, 124, 211}

func greet(name string) {
	fmt.Printf("Hello, %v\n", name)
}
