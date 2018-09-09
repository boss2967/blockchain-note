package main

import "fmt"
//切片不能通过 = 互相赋值
func main() {
	src:=[]int{5, 6, 7, 8, 9}
	outSlice:=deleteElementByPosition(src,2)
	fmt.Println(outSlice)
}
func deleteElementByPosition(src []int,index int)[]int {
	copy(src[index:],src[index+1:])//两个切片之间互相赋值要用copy
	return src[:len(src)-1]
}
