package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{1, 2, 3, 4, 5, 6}
	arr3 := []int{1, 2, 3, 4, 5, 6}
	arr4 := []int{1, 2, 3, 4, 5, 6}
	arr5 := []int{1, 2, 3, 4, 5, 6}
	//删除
	Delete(arr1)
	fmt.Println("删除", arr1) //[1 2 3 4 5 6]
	//修改
	Modify(arr2)
	fmt.Println("修改", arr2) //[0 2 3 4 5 6]
	//追加
	Append(arr3)
	fmt.Println("追加", arr3) //[1 2 3 4 5 6]
	//先删除第一个再追加
	DeleteFirstThenAppend(arr4)
	fmt.Println("先删除第一个再追加", arr4) //[1 2 3 4 5 6]
	//先删除非第一个再追加
	DeleteNotFirstThenAppend(arr5)
	fmt.Println("先删除非第一个再追加", arr5) //[1 2 3 4 5 10]
}
func DeleteNotFirstThenAppend(arr []int) {
	arr = arr[:len(arr)-1]
	arr = append(arr, 10)
}
func DeleteFirstThenAppend(arr []int) {
	arr = arr[1:]
	arr = append(arr, 20)
}

func Append(arr []int) {
	arr = append(arr, 10)
}
func Modify(arr []int) {
	arr[0] = 0
}
func Delete(arr []int) {
	arr = arr[:len(arr)-1]
}
