package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string //和用户绑定的channel
	conn net.Conn    //当前用户和客户端通信的链接句柄

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户的上线业务
func (user *User) Online() {

	//用户上线，创建一个user实例

	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.BroadCast(user, "已上线")
}

// 用户的下线业务
func (user *User) Offline() {
	//用户下线，删除

	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.BroadCast(user, "下线")
}

//s

func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	fmt.Println(msg)
	if msg == "who" {
		//查询当前在线用户
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + u.Addr + "]" + u.Name + ":" + "在线\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else {
		//将用户发送的消息进行广播
		user.server.BroadCast(user, msg)
	}

}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C //从用户中取值

		user.conn.Write([]byte(msg + "\n"))
	}
}
