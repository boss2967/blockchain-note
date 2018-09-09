package main

import (
	"math/rand"
	"time"
	"fmt"
)

//这种写法有两个问题：
//1，有可能一个消费者在读取了channel但是并没有处理完，下一个生产者写入，消费者又开始读取处理了，数据有可能有误
//2，如果消费者比较多，或者比较快，需要不断的检查是否有新数据，会造成cpu资源浪费
var sum int

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	for i := 0; i < 5; i++ {
		go producter1(ch)
	}
	go customer1(ch)
	customer1(ch)
	<-quit
}
func customer1(ch <-chan int) {
	for n := range ch {
		sum += n
		fmt.Println("消费者", sum)
	}
}
func producter1(ch chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		fmt.Println("生产者", n)
		ch <- n
	}
}
