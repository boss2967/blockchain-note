package main

import "fmt"

/*
	可以指定索引去初始化
*/
func main() {
	//0 make 只能创建slice map channel
	slice0 := make([]int, 5, 5)
	//1 如果不指定那么容量等于长度
	slice1 := make([]int, 5)
	//2
	slice2 := []int{1, 2, 3, 4}
	//3
	slice3 := []string{1: "一月", 2: "二月"}

	fmt.Println(len(slice3))                    //3
	fmt.Println(slice0, slice1, slice2, slice3) //[0 0 0 0 0] [0 0 0 0 0] [1 2 3 4] [ 一月 二月]
}
