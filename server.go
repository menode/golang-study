package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port string

	//在线用户的列表
	OnlineMap map[string]*User
	//由于map是全局的，需要加一个读写锁
	mapLock sync.RWMutex
	//消息广播的channel
	Message chan string
}

// 广播消息的方法 服务端循环给所有在线用户发送消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

// 监听消息广播channel的goroutine，一旦有消息就发送给全部在线用户
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message //不断尝试从message channel中获取数据
		//将消息发送给全部在线用户
		s.mapLock.Lock()
		//将msg发送给全部在线用户
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}
func (s *Server) handleRequest(conn net.Conn) {
	// fmt.Println("链接成功")
	user := NewUser(conn)
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()
	//广播当前用户上线消息
	s.BroadCast(user, "已上线")
	//发送完消息前先堵塞Handle 保证goroutine不会退出
	select {}
	//当前链接的用户不断的监听channel

	// buffer := make([]byte, 1024)
	// _, err := conn.Read(buffer)
	// if err != nil {
	// 	fmt.Println("Error reading:", err.Error())
	// }
	// conn.Write([]byte("Message received."))
	// conn.Close()
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.IP+":"+s.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close() //为了防止遗忘关闭，使用defer关闭

	fmt.Println("Listening on " + s.IP + ":" + s.Port)
	//监听Message广播消息channel
	go s.ListenMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go s.handleRequest(conn)
	}
}

func NewServer(ip, port string) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}
