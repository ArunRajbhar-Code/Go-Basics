package main

import "time"

func task(i int) {
	println("running task", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i) //here go will run this functions ten times parrallely

	}
	time.Sleep(time.Second * 2)
}
