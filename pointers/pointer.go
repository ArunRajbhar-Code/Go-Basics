package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c

}

func main() {
	var a int = 7
	var b int = 6
	x := 34
	p := &x //address poiter of variable x
	fmt.Println(x, p)
	fmt.Println("Before swaping", "\na= ", a, "\nb= ", b)
	swap(&a, &b)
	fmt.Println("After swap", "\na= ", a, "\nb= ", b)
}
