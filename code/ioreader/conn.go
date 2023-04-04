package main

import (
	"fmt"
	"io"
	"net"
)

// 处理客户端请求
func handleConn(con *net.TCPConn) {
	// 读取客户端发送的数据
	data, err := io.ReadAll(con)
	// 如果出现错误，则关闭连接，否则不关闭连接
	if err != nil {
		fmt.Printf("get client data error:%v\n", err)
		//
		con.Close()
		return
	}
	fmt.Printf("client send data=%s\n", data)
}
func main() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8888,
	}
	// 监听服务器, 返回listen 对象和err
	listen, err := net.ListenTCP("tcp", addr)
	// 关闭服务器
	defer listen.Close()
	if err != nil {
		fmt.Printf("listen error:%v\n", err)
		return
	}
	for {
		// 接收客户端连接请求，成功返回Conn 对象
		// err 表示连接错误
		conn, cErr := listen.AcceptTCP()
		if cErr != nil {
			fmt.Printf("connect error:%v\n", cErr)
		}
		go handleConn(conn) // 开心gorutinues
	}

}
