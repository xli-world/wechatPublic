package main

import "fmt"

func main() {
	p := new(struct{})
	q := new(struct{})
	fmt.Println(p == q)
}
