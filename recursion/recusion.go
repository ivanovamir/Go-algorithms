package main

import (
	"fmt"
	"time"
)

func main() {
	ReqFunc(5)
}

func ReqFunc(i int) {
	if i <= 0 {
		return
	} else {
		time.Sleep(1 * time.Second)
		fmt.Printf("time remaining: %d\n", i)
		ReqFunc(i - 1)
	}
}
