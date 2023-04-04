package main

import "net"

func main() {
	// 创建listen
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
}
