package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	var x, y int
	go func() {
		x = 1
		fmt.Print("y:", y, " ")
	}()
	go func() {
		y = 1
		fmt.Print("x:", x, " ")
	}()
	time.Sleep(time.Second)
}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 100; i++ {
		go func() {
			test()
			fmt.Println()
		}()
	}
}
