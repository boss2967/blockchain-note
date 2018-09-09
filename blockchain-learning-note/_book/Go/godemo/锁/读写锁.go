package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var count int //定义一个全局变量
var rwMutex sync.RWMutex

func main() {
	quit := make(chan bool)
	for i := 0; i < 5; i++ {
		go write()
	}
	for i := 0; i < 5; i++ {
		go read()
	}
	<-quit
}
func read() {
	for {
		rwMutex.RLock()
		fmt.Println("读取", count)
		rwMutex.RUnlock()
	}
}
func write() {
	for {
		rwMutex.Lock()
		rand.Seed(time.Now().UnixNano())
		count = rand.Intn(10)
		fmt.Println("写入", count)
		rwMutex.Unlock()
	}
}
