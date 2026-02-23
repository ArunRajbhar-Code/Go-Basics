package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o order) getID() string {
	return o.id
}
func (o order) getAmount() float32 {
	return o.amount
}
func (o order) getStatus() string {
	return o.status
}

func (o *order) setAmount(amount float32) {
	o.amount = amount

}
func (o *order) setID(id string) {
	o.id = id
}
func (o *order) setStatus(status string) {
	o.status = status
}

func main() {
	order1 := order{
		id:        "1at",
		amount:    23.89,
		status:    "recieved",
		createdAt: time.Now(),
	}
	order2 := order{
		id:        "1atew",
		amount:    27.89,
		status:    "recieved",
		createdAt: time.Now(),
	}
	order2.setAmount(54.90)

	fmt.Println(order1)
	fmt.Println(order2)

}
