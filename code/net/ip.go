package main

import (
	"fmt"
	"net"
)

func main() {
	// ip 字符串
	ip := "127.0.0.1"
	// 解析字符串返回IP
	Ip := net.ParseIP(ip)
	fmt.Println(Ip)
	// 不是ip ，解析失败返回nil
	fmt.Println(net.ParseIP("127.0.0"))
}
