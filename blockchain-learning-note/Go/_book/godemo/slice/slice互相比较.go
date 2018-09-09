package main

import (
	"fmt"
)

//1，数组，结构体之间可以 互相比较 ，map不可以互相比较
//2，slice 之间不能互相比较，只能和nil 比较。结构体包含了map或slice也不可以比较
//slice can only be compared to nil
//第一个原因：一个 slice 的元素是间接引用的，一个 slice甚至可以包含自身。虽然有很多办法处理这种情形，但是没有一个是简单有效的。
type stru1 struct {
}

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1 == arr2)//true
	st1 := stru1{}
	st2 := stru1{}
	fmt.Println(st1 == st2) //true
	//s1 := make([]int, 10)
	//s2 := make([]int, 10)
	//s2 = append(s2, 10)
	//fmt.Println(s1 == s2) //error slice can only be compared to nil
}
