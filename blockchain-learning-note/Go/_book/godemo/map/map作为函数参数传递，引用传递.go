package main

import "fmt"

//可以改变值

func main() {
	m:=map[int]string{
		1:"yi",
		2:"er",
	}
	deleteMap(m)
	fmt.Println(m)//[2:er]
	addMap(m)
	fmt.Println(m)//[3:san 2:er]
	changeMap(m)
	fmt.Println(m)//[3:san 2:second]
}
func changeMap(m map[int]string) {
	m[2]="second"
}
func addMap(m map[int]string) {
	m[3]="san"
}
func deleteMap(m map[int]string) {
	delete(m,1)
}
