package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN
	fmt.Println(1/z == math.Inf(1))    // true
	fmt.Println(z/z == math.NaN())     // false
	fmt.Println(math.IsNaN(z / z))     // true
	var m int64
	fmt.Println(m, -m, 1/m, -1/m, m/m) // 0 -0 +Inf -Inf NaN
}
