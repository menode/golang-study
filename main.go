package main

import (
	"fmt"
)

func main() {
	server := NewServer("localhost", "8080")
	server.Start()
	fmt.Println(server)
}
