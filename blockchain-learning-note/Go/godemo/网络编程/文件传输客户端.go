package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

var path = "/Users/zhengjunling/Downloads/picture/linux-kernel-map.jpg"

func main() {
	fileInfo,_:=os.Stat(path)
	//和服务器创建连接
	conn, err := net.Dial("tcp", "127.0.0.1:8686")
	defer conn.Close()
	ControErr(err, "net dial err:")
	sendFileNameToServer(conn, fileInfo.Name())
	reciveFromServer(conn)
}
func reciveFromServer(conn net.Conn) {
	b := make([]byte, 4096)
	var result string
	n, err := conn.Read(b)
	if err != nil {
		return
	}
	result = string(b[:n])

	fmt.Println("result=", result)
	if result == "ok" {
		sendFileToServe(conn)
	}
}
func sendFileToServe(conn net.Conn) {
	f, err := os.Open(path)
	defer f.Close()
	ControErr(err, "os open err:")
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		conn.Write(buf[:n])
		if err == io.EOF {
			fmt.Println("发送完成")
			break
		}
		if err != nil {
			return
		}
	}
}
func sendFileNameToServer(conn net.Conn, fileName string) {
	//给服务器发送消息
	_, err := conn.Write([]byte(fileName))
	ControErr(err, "conn write err:")
}
func ControErr(err error, errMessage string) {
	if err != nil {
		fmt.Println(errMessage, err)
		os.Exit(1)
	}
}
