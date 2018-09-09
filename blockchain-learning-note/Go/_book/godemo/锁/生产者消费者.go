package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var quit = make(chan bool)

var con sync.Cond

func main() {
	ch := make(chan int,3)
	con.L = new(sync.Mutex)
	for i := 0; i < 15; i++ {
		go producter(ch, i)
	}
	for i := 0; i < 5; i++ {
		go customer(ch,i)
	}
	<- quit
	//for{
	//	;
	//}
}
func customer(ch <-chan int,i int) {
	for{
		con.L.Lock()
		for len(ch)==0{
			con.Wait()
		}
		num:=<-ch
		fmt.Println("线程 ", i, "消费者 消费 ", "---", num)
		con.L.Unlock()
		con.Signal()
		time.Sleep(time.Millisecond * 500)
	}
}
func producter(ch chan<- int, i int) {
	if i == -1 {
		quit<- true
	}
	rand.Seed(time.Now().UnixNano())
	for {
		con.L.Lock()
		if len(ch)==3{
			con.Wait()
		}
		num := rand.Intn(10)
		fmt.Println("线程 ", i, "生产者 生产 ", "---", num)
		ch<-num
		con.L.Unlock()
		con.Signal()
		time.Sleep(time.Millisecond * 500)
	}
}
