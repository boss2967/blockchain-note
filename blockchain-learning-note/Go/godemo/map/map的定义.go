package main

import (
	"fmt"
	"time"
)

//-1，map不可以互相比较
//0，map的特性键不可重复，无序，在 Go 语言中，一个map 就是一个哈希表的引用
//1，map的键不能是function，map，slice 值可以是任意类型
//2，map可以指定容量 但是不能使用cap(map)，只能使用len(map)
//3，可以创建空的map   map[string]int{}
//4，可以创建并直接指定map的元素
//5，先指定再添加元素
//6，通过下标判断是否有该键，第一个是值，第二个是bool 		v,ok:=m["f"]
//7，delete 没有返回值，即使查找的元素不在map中也没问题 返回nil
//8，禁止对map元素取值，map 可能随着元素数量的增长而重新分配 更大的内存空间，仅而可能寻致之前的地址无效。
func main() {
	arr:=[1]int{3}
	m:=make(map[[1]int]interface{})
	m[arr]="v"
	fmt.Println(m)
	m1 := make(map[int]interface{})
	//m11 := make(map[int]map[int]int)
	//m12 := make(map[int]map[int]int)
	//fmt.Println(m11==m12)
	//m1 := make(map[map[int]int]map[int]int)//err
	fmt.Println(m1, "map的键不能是function，map，slice 值可以是任意类型")
	m2 := make(map[int]map[int]interface{}, 10)
	fmt.Println(m2, "len=", len(m2), "map可以指定容量 但是不能使用cap(map)，只能使用len(map)")
	m2[1] = m1
	fmt.Println(m2, "len=", len(m2), "map可以指定容量 但是不能使用cap(map)，只能使用len(map)")
	//cap(m2)//err
	m3 := map[int]string{}
	time.Now()
	fmt.Println(m3, "可以创建空的map   map[string]int{}")
	var f float64
	m4 := map[interface{}]interface{}{
		"h":   1,
		true:  2,
		f / f: 2,
	}
	fmt.Println(m4, "可以创建并直接指定map的元素", m4[f/f], "尽量不用使用float作为键")
	m5 := make(map[interface{}]interface{})
	m5["haha"] = 1
	m5[true] = false
	v, ok := m5["dfads"]
	fmt.Println(m5, "v:", v, "ok:", ok, "通过下标判断是否有该键，第一个是值，第二个是bool")
	delete(m5, "haha")
	fmt.Println(m5[false])
	//fmt.Println(&m5[false])//err

}
