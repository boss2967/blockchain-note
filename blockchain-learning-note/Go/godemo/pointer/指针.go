package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ii [4]int = [4]int{ 11, 22, 33, 44 }

	var p0 * int          = &ii[0]               // p0 point to first element
	var p1 unsafe.Pointer = unsafe.Pointer(p0)   // convert (* int) to unsafe.Pointer
	var p2 uintptr        = uintptr(p1)          // convert unsafe.Pointer to uintptr
	p2 += 8                                      // computing uintptr with plus 8, i.e, the next element address
	var p3 unsafe.Pointer = unsafe.Pointer(p2)   // convert uintptr back to unsafe.Pointer
	var p4 * int64        = (* int64)(p3)        // convert unsafe.Pointer to another type pointer, (* int64)

	fmt.Printf("*p0=%d,*p1=%d,*p2=%d,*p3=%d,*p4=%d\n", *p0, p1, p2, p3, *p4)
}