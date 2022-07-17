package main

import (
	"log"
	"os"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...这里是一些处理...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func traverse(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // 注意：可能会用尽文件描述符
		// ...处理文件f...
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...处理文件f...
	return nil
}

func main() {
	bigSlowOperation()
}
