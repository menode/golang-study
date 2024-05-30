package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string   //服务器IP
	ServerPort int      //服务器端口
	Name       string   //用户名
	conn       net.Conn //链接
	flag       int      //当前客户端模式
}

func NewClient(serverIP string, serverPort int) *Client {
	//创建客户端
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}
	//链接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn
	//返回对象
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>> 请输入合法范围内的数字...")
		return false
	}
}

func (client *Client) DealResponse() {
	//一旦client.conn有数据，就会调用这个函数
	//并且不断的将数据发送到客户端
	io.Copy(os.Stdout, client.conn)

}

func (client *Client) PublicChat() {
	fmt.Println(">>>>>> 请输入聊天内容，exit退出:")
	var chatMsg string
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		//发给服务器
		_, err := client.conn.Write([]byte(chatMsg + "\n"))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			break
		}
		chatMsg = ""
		fmt.Println(">>>>>> 请输入聊天内容，exit退出:")
		fmt.Scanln(&chatMsg)

	}

}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

}

// 私聊模式
func (client *Client) PrivateChat() {
	client.SelectUsers()
	fmt.Println(">>>>>> 请输入聊天对象[用户名] ,exit退出")
	var remoteName string
	fmt.Scanln(&remoteName)
	fmt.Println(">>>>>> 请输入聊天内容")
	var chatMsg string
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		//消息格式：to|张三|消息内容
		sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
		_, err := client.conn.Write([]byte(sendMsg))
		if err != nil {

			fmt.Println("conn.Write err:", err)
			break
		}
		chatMsg = ""
		fmt.Println(">>>>>> 请输入聊天内容，exit退出:")
		fmt.Scanln(&chatMsg)
	}
	client.SelectUsers()
	fmt.Println(">>>>>> 请输入聊天对象[用户名] ,exit退出")
	fmt.Scanln(&remoteName)
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>> 请输入用户名:")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true

}

func (client *Client) Run() {
	for client.flag != 0 {

		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			client.PublicChat()
			break
		case 2:
			client.PrivateChat()
			break
		case 3:
			client.UpdateName()
			break
		case 0:
			fmt.Println(">>>>>> 退出...")
			break
		}
	}

}

var ServerIp string
var ServerPort int

func init() {
	fmt.Println(">>>>>> init...")
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "setting server ip")
	flag.IntVar(&ServerPort, "port", 8080, "setting server port")
}

func main() {
	//命令行解析
	flag.Parse()
	//客户端链接服务器

	client := NewClient(ServerIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>>> connect server error...")
		return
	}
	go client.DealResponse()

	fmt.Println(">>>>>> connect server success...")

	//保持客户端不断开
	client.Run()
}
