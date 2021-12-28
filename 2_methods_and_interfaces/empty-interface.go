package main

import "fmt"

func main() {
	var i interface{} // The interface type that specifies zero methods is known as the empty interface
	describe(i)

	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	i = 42
	describe(i)

	i = "hello"
	describe(i)

	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
}

func describe(i interface{}) {
	fmt.Printf("(%v, %[1]T)\n", i)
}
