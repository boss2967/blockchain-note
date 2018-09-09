package main

import (
	"unsafe"
	"fmt"
)

/*
Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：
1) 任意类型的指针可以转换为一个Pointer类型值
2) 一个Pointer类型值可以转换为任意类型的指针
3) 一个uintptr类型值可以转换为一个Pointer类型值
4) 一个Pointer类型值可以转换为一个uintptr类型值
*/
func main() {
	var ptrI *int
	ptrI = new(int)
	*ptrI = 1
	fmt.Println(*(*int)(unsafe.Pointer(ptrI)))
	fmt.Println(*(*int64)(unsafe.Pointer(ptrI)))
}
