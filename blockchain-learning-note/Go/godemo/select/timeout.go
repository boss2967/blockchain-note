package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(time.Second * 3):
				quit <- true
			}
		}
	}()
	ch <- 666
	<-quit
}
