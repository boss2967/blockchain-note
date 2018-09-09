package main

import (
	"net"
	"io"
	"fmt"
	"os"
)

var fileName string

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8686")
	defer listener.Close()
	ControErr(err, "net listen err:")
	conn, err := listener.Accept()
	defer conn.Close()
	ControErr(err, "listener accept err:")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fileName = string(buf[:n])
	addr := conn.RemoteAddr()
	fmt.Printf("[%s]：%s", addr, fileName)
	sendOkToClient(conn)
	reciveFromClient(conn)
}
func reciveFromClient(conn net.Conn) {
	buf := make([]byte, 4096)
	f, err := os.Create(fileName)
	ControErr(err, "os create err:")
	for {
		n, err := conn.Read(buf)
		f.Write(buf[:n])
		if err == io.EOF {
			fmt.Println("接收完成")
			break
		}
		if err != nil {
			return
		}
	}
}
func sendOkToClient(conn net.Conn) {
	conn.Write([]byte("ok"))
}
func ControErr(err error, errMessage string) {
	if err != nil {
		fmt.Println(errMessage, err)
		os.Exit(1)
	}
}
