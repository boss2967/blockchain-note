package main

import "fmt"

/*
	1，append第二个参数使用切片，
*/
func main() {
	ss := make([]int, 0)
	s := make([][]int, 0)
	ss = append(ss, 1)
	s = append(s, ss)
	fmt.Println(s)

	c:=[]int{1}
	cc:=make([]int,0)
	cc = append(cc,c...)
	fmt.Println(cc)
}
