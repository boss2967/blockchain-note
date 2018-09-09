package main

import (
	"math/rand"
	"fmt"
	"sync"
)
var rwMutex sync.RWMutex
var num int
func main() {
	for i:=0;i<5 ;i++  {
		go read(i)
	}
	for i:=0;i<5 ;i++  {
		go write(i)
	}
	for{
		;
	}
}
func write(i int) {
	rwMutex.Lock()
	fmt.Printf("写 goroutine %d 正在写取数据...\n", i)
	defer rwMutex.Unlock()
	n := rand.Intn(1000)
	num = n
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n",i,n)
}
func read(i int) {
	fmt.Printf("读 goroutine %d 正在读取数据...\n", i)
	rwMutex.RLock()
	rwMutex.RUnlock()
	fmt.Printf("读 goroutine %d 读数据结束，读取新值 %d\n",i,num)
}
