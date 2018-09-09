package main

import (
	"strings"
	"fmt"
)

func main() {
	str:="I love my work and I love my family too"
	m:=make(map[string]int)
	wcFunc(str,m)
	fmt.Println(m)
	m1:=make(map[string]int)
	wcFunc1(str,m1)
	fmt.Println(m1)
}
func wcFunc1(s string, m1 map[string]int) {
	strs:= strings.Fields(s)
	for i:=0;i<len(strs);i++ {
		if _,ok:=m1[strs[i]];ok{
			m1[strs[i]]++
		}else {
			m1[strs[i]]=1
		}
	}
}
func wcFunc(s string, m map[string]int) {
	strs:= strings.Fields(s)
	for i:=0;i<len(strs);i++{
		m[strs[i]]++
	}
}
