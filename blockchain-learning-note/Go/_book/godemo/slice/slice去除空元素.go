package main

import "fmt"

func main() {
	strs:= []string{"red", "", "black", "", "", "pink", "blue"}
	outStrs:=deleteEmptyString(strs)
	fmt.Println(outStrs)
}
func deleteEmptyString(strs []string)[]string {
	outStrs:=make([]string,0)
	for _,v:=range strs{
		if v!="" {
			outStrs = append(outStrs,v)
		}
	}
	return outStrs
}
