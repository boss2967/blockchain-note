package main

import "fmt"

//map   _,ok:= m[V]
func main() {

	srcArr:=[]string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	dst:=deleteRepeatElement(srcArr)
	fmt.Println(dst)
	dst2:=deleteEmptyStringByMap(srcArr)
	fmt.Println(dst2)
}
func deleteEmptyStringByMap(src []string)(dst []string) {
	m:=make(map[string]int)
	for _,v:= range src{
		m[v]=0
	}
	for k:=range m{
		dst= append(dst,k)
	}
	return
}
func deleteRepeatElement(src []string)(dst[]string) {
	for _,v:=range src{
		flag:=false
		//先判断输出的slice里面有没有
		for _,str:=range dst{
			//设置一个flag
			if str==v {
				flag=true
			}
		}
		//没有就append
		if !flag {
			dst = append(dst,v)
		}
	}
	return
}
