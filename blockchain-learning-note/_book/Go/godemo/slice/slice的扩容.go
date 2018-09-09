package main

import "fmt"

/*
	append会只能的将slice的底层数组容量增加，
	通常以两倍的容量重新分配底层数组(在1024以下时)，并复制原来内容,
	因此使用append给切片扩容时，切片地址可能发生改变
*/

func main() {
	slice := make([]int,0)
	for i:=0;i<20 ;i++  {
		slice = append(slice,i)
		fmt.Println(cap(slice))
	}
}
