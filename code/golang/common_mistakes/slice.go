package main

import "fmt"

func main() {
	var months []string = []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
	summer := months[3:6]
	fmt.Println(summer[:5]) // [7月 8月 9月 10月 11月]
	//fmt.Println(summer[5])  // panic: runtime error: index out of range [5] with length 3
	var b = []string{"123", "234", "345", "456"}
	var slowLength = make([]string, 0, len(b))
	for index := range b {
		slowLength[index] = b[index]
	}

	slowLength[0] = "123"
	fmt.Println(len(slowLength), slowLength[len(slowLength)])
}
