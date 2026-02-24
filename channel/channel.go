package main

import (
	"fmt"
	"time"
)

func printNum(numChan chan int) {
	fmt.Println(<-numChan)

}

func main() {
	numChan := make(chan int)
	go printNum(numChan)
	numChan <- 5
	time.Sleep(time.Second * 2)

	// messageChan := make(chan string)
	// messageChan <- "hello channel"
	// msg := <-messageChan
	// fmt.Println(msg)
}
