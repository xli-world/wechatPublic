package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		fmt.Println("%s", <-c)
	}()
	c <- "123"
	time.Sleep(1 * time.Second)
}
