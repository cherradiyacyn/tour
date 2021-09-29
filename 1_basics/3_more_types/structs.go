package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Printf("Type : %T\tValue : %[1]v\n", Vertex{1, 2})
}
