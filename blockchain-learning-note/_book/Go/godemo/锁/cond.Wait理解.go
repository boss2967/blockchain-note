package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i int
	rand.Seed(time.Now().UnixNano())
	for {
		i = rand.Intn(11)
		//for {
		//	time.Sleep(time.Second)
		//	fmt.Println(i)
		//	for i == 10 {
		//		fmt.Println("i=10!!!")
		//	}
		//}
		//fmt.Println("finish")

		for 10 == i {
			fmt.Println()
		}
	}

}
