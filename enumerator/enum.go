package main

import "fmt"

type Status int

const ( // enumarator implementation
	Pending Status = iota
	Processing
	Completed
	Cancelled
)

func main() {
	var orderStatus Status = Processing
	fmt.Println(orderStatus) // Output: 1
}
