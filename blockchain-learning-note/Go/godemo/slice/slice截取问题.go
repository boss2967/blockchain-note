package main

import "fmt"

/*
	一个切片由指针（首元素），长度，容量 组成
	截取的时候左闭右开
	s[low:high:max]
	len = high-low
	cap = max-low
*/

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	s1 := s[:3]
	fmt.Println(s1)      //[1 2 3]
	fmt.Println(cap(s1)) //6
	s2 := s1[:4]         //根据首元素的地址找到底层数组的起始位置，向后截取
	fmt.Println(s2)      //[1 2 3 4]
	//s3:= s1[4] 		 //error 下标 可能是根据s1长度遍历取出来的
	//fmt.Println(s3)
	s11 := s[3:4]
	fmt.Println(s11)      //[4]
	fmt.Println(cap(s11)) //3 由于没有指定s1的cap，那么cap就为s1的首元素到原切片末尾的长度
}
