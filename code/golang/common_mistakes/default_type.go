package main

import "fmt"

func main() {
	i := 0
	r := `\000`
	f := 0.0
	c := 0i
	fmt.Printf("%T %T %T %T", i, r, f, c) // int string float64 complex128
}
