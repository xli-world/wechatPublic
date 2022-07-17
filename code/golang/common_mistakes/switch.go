package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	filename := "test"
	f, err := os.Open(filename) // 声明了两个新变量f，err
	_ = f
	fmt.Println(err)                                 // open test: no such file or directory
	if m, err := os.Readlink(filename); err != nil { // err为新声明的局部变量
		fmt.Println(err) // readlink test: no such file or directory
		_ = m
	}
	fmt.Println(err)                                   // open test: no such file or directory
	for n, err := os.Readlink(filename); err != nil; { // err为新声明的局部变量
		fmt.Println(err) // readlink test: no such file or directory
		_ = n
		break
	}
	fmt.Println(err)                                  // open test: no such file or directory
	p, err := "new variable", errors.New("new error") // 未声明新变量err，而是将errors.New("new error")赋值给了之前的err变量
	fmt.Println(err)                                  // new error
	_ = p
}
