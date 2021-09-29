package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // p holds the memory address of the value 42
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer
	fmt.Println(i)  // i holds the value 21

	p = &j
	*p = *p / 37
	fmt.Println(*p) // *p : dereferencing
}
