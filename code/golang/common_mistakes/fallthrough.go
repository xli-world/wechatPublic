package main

import "fmt"

func main() {
	var i = 5
	switch {
	case i <= 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case i <= 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case i <= 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case i <= 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case i <= 3:
		fmt.Println("The integer was <= 3")
	default:
		fmt.Println("default case")
	}
}
