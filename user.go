package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string //和用户绑定的channel
	conn net.Conn    //当前用户和客户端通信的链接句柄
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C //从用户中取值

		user.conn.Write([]byte(msg + "\n"))
	}
}

// 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}
