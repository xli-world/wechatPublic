package main

import "fmt"

func main() {
	medals := []string{"gold", "silver", "bronze"}
	var length uint32 = 3
	for i := length - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}
