package main

import (
	"strconv"
	"fmt"
	"net/http"
	"io"
	"github.com/json-iterator/go"
	"os"
)

var url = "https://www.douyu.com/gapi/rkc/directory/2_201/"
var dstUrl = "/Users/zhengjunling/Downloads/awesomeProject/src/douyu/yanzhi/"
//
//var url = "https://www.douyu.com/gapi/rkc/directory/2_1/"
//var dstUrl = "/Users/zhengjunling/Downloads/awesomeProject/src/douyu/lol/"

func main() {
	ch := make(chan int)
	const TOTAL = 14
	for i := 0; i < TOTAL; i++ {
		go getPageUrl(url+strconv.Itoa(i), i, ch)
	}
	for i := 0; i < TOTAL; i++ {
		fmt.Printf("第%d页爬去完成\n", <-ch)
	}
}
func getPageUrl(url string, i int, ch chan<- int) {
	fmt.Println("url:", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	defer response.Body.Close()
	buf := make([]byte, 4096)
	var content string
	for {
		n, err1 := response.Body.Read(buf)
		content += string(buf[:n])
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			fmt.Println("err1:", err1)
			return
		}
	}
	//fmt.Println(content)
	val := []byte(content)
	for i := 0; ; i++ {
		if jsoniter.Get(val, "data", "rl", i, "rs1").ToString() == "" {
			return
		}
		fmt.Println(jsoniter.Get(val, "data", "rl", i, "rs1").ToString())
		pictureUrl := jsoniter.Get(val, "data", "rl", i, "rs1").ToString()
		name := jsoniter.Get(val, "data", "rl", i, "nn").ToString()
		go fetchThePictrue(name, pictureUrl)
	}
	ch <- i
}
func fetchThePictrue(name string, pictureUrl string) {
	response,err:=http.Get(pictureUrl)
	if err!=nil{
		fmt.Println("http get err:",err)
		return
	}
	defer response.Body.Close()
	f,err2:=os.Create(dstUrl+name+".jpg")
	if err2!=nil{
		fmt.Println("os create err:",err2)
		return
	}
	defer f.Close()
	buf:=make([]byte,4096)
	for {
		n, err1 := response.Body.Read(buf)
		f.Write(buf[:n])
		if err1==io.EOF{
			fmt.Println("end")
			break
		}
		if err1!=nil{
			fmt.Println("response body read err:",err1)
			return
		}
	}
}
