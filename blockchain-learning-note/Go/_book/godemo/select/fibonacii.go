package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
		quit <- true
	}()
	getFibonaci(ch, quit)
}
func getFibonaci(ch chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
