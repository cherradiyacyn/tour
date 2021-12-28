package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type assertion provides access to an interface value's underlying concrete value.
	s := i.(string)
	// This statement asserts that the interface value i holds the concrete type string and assigns the underlying string value to the variable s.
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
