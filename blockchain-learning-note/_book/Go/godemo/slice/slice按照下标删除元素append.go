package main

import "fmt"
//利用append 的第二个不定长参数
func main() {
	arr := []int{5, 6, 7, 8, 9}
	outSlice := deleteElementAppend(arr, 2)
	fmt.Println(outSlice)
}
func deleteElementAppend(src []int, i int) []int {
	return append(src[:i], src[i+1:]...)//append第二个形参时不定长参数
}
