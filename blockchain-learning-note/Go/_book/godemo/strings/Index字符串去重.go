package main

import (
	"strings"
	"fmt"
)

func main() {
	postion:=strings.Index("abcdvc","c")
	fmt.Println(postion)
	str:="aaaaadddsfgrge"
	resultArr:=[]byte{}
	for i:=0;i<len(str) ;i++  {
		if strings.Index((str),string(str[i]))==i{
			resultArr=append(resultArr, str[i])
		}
	}
	fmt.Println(string(resultArr))

}
