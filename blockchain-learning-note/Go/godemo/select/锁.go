package main

import (
	"fmt"
	"sync"
)
var mutext sync.Mutex
func main() {
	go printer1()
	go printer2()
	for{
		;
	}
}
func printer1(){
	printer("1111111111111111111111111111")
}
func printer2(){
	printer("2222222222222222222222222222")
}
func printer(s string){
	mutext.Lock()
	defer mutext.Unlock()
	for _,data:=range s{
		fmt.Println(string(data))
		//time.Sleep(time.Second) 放大go程竞争效果
	}
	fmt.Println()
}