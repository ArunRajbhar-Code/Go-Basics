package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter
}
type stripe struct {
}

func (s stripe) pay(amount float32) {
	println(amount, "amount has been payed through strip")
}

type rozerPay struct {
}

func (s rozerPay) pay(amount float32) {
	println(amount, "amount has been payed through razerpay")
}

func main() {
	fmt.Println("interface")
	// Using Stripe
	s := stripe{}
	p1 := payment{
		gateway: s,
	}
	p1.gateway.pay(34.90)

	// Using RazorPay
	r := rozerPay{}
	p2 := payment{
		gateway: r,
	}
	p2.gateway.pay(50.00)

}
