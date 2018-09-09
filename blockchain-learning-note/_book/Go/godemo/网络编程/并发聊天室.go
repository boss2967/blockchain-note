package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

type Client struct {
	C       chan string
	Address string
	Name    string
	conn    net.Conn
}

var Message = make(chan string)
var userMap = make(map[string]*Client)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	HandleErr(err, "net listener err:")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		HandleErr(err, "listener Accept err:")
		defer conn.Close()
		go handleMsg(Message)
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	C := make(chan string)
	address := conn.RemoteAddr()
	message := createMsg(address.String(), "上线了")
	fmt.Println(message)
	Message <- message
	user := Client{C, address.String(), address.String(), conn}
	userMap[address.String()] = &user
	//	把新登录的用户添加至全局用户map中
	//	广播给所有的用户
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		HandleErr(err, "conn read err:")
		content := string(buf[:n-1])
		fmt.Println("content:", content)
		message = createMsg(address.String(), content)
		if strings.HasPrefix(content, "rename") {
			fmt.Println("rename|")
			//rename
			userMap[address.String()].Name = strings.Split(content, "|")[1]
			go writeToClient(&user)
			user.C <- "rename success !!!\n"
			continue
		}
		if content == "showusers" {
			var showUserMsg string
			fmt.Println("showusers")
			for _, v := range userMap {
				showUserMsg += "[" + v.Name + "]\n"
			}
			go writeToClient(&user)
			user.C <- showUserMsg
			continue
		}
		Message <- message
	}
}
func handleMsg(Msg chan string) {
	for {
		fmt.Println("handleMsg")
		message := <-Msg
		for _, v := range userMap {
			fmt.Println("range")
			go writeToClient(v)
			v.C <- message
		}
	}
}
func writeToClient(client *Client) {
	msg := <-client.C
	client.conn.Write([]byte(msg))
	fmt.Println("writeToClient")
}
func createMsg(addr string, s string) string {
	return "[" + addr + "]" + s + "\n"
}
func HandleErr(err error, errMessage string) {
	if err != nil {
		fmt.Println(errMessage, err)
		os.Exit(1)
	}
}
