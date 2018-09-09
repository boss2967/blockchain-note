package main

import "fmt"
/*
	目标元素用末尾元素替换
*/
func main() {
	arr:=[]int{5, 6, 7, 8, 9}
	outSlice:=deleteElementByPositionNoOrder(arr,2)
	fmt.Println(outSlice)
}
func deleteElementByPositionNoOrder(src []int,index int)[]int {
	src[index]=src[len(src)-1]
	return src[:len(src)-1]
}
