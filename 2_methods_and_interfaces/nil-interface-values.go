package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I // A nil interface value holds neither value nor concrete type.

	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %[1]T)\n", i)
}
