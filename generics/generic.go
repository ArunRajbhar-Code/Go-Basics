package main

import "fmt"

// T is a type parameter
func Print[T any](value T) {
	fmt.Println(value)
}

type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func main() {
	Print(10)
	Print("Hello")
	Print(3.14)
	intBox := Box[int]{value: 100}
	strBox := Box[string]{value: "Go"}

	fmt.Println(intBox.Get())
	fmt.Println(strBox.Get())
}
