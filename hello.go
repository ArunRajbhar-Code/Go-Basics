package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	var a int = 8
	var b int = 9
	fmt.Println(a + b)
	if a > b {
		fmt.Println("greater")
	} else {
		fmt.Println("lower")
	}
}
