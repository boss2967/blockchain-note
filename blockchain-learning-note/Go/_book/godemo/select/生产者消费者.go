package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)
var cond sync.Cond
func main() {
	rand.Seed(time.Now().UnixNano())
	//quit:=make(chan bool)
	product:=make(chan int,3)
	cond.L=new(sync.Mutex)
	for i:=0;i<100 ;i++  {
		go producter(product,i)
	}
	for i:=0;i<100 ;i++  {
		go customer(product,i)
	}
	//<-quit
	for {
		fmt.Println("for")
	}
}
func producter(out chan<- int,i int){
	for {
		cond.L.Lock()
		for len(out)==3{
			cond.Wait()
		}
		num := rand.Intn(1000)
		out<- num
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", i, num, len(out))
		cond.L.Unlock()
		cond.Signal()
	}
}
func customer(in <-chan int,i int){
	for{
		cond.L.Lock()
		for len(in)==0{
			cond.Wait()
		}
		num:=<- in
		fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", i, num, len(in))
		cond.L.Unlock()
		cond.Signal()
	}
}
