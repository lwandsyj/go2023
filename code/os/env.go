package main

import (
	"fmt"
	"os"
)

// 使用os 中的 env 相关的函数设置和获取环境变量
func main() {
	// 更加环境变量名称获取环境变量的值
	node_env := os.Getenv("NODE_ENV")
	//
	url := os.Getenv("url")

	fmt.Printf("node_env= %s\n", node_env)

	fmt.Printf("url=%s\n", url)

	// os.Setenv(key,value string) error 设置环境变量
	os.Setenv("url", "http://www.baidu.com")
	// 获取环境变量
	url = os.Getenv("url")

	fmt.Printf("new url=%s\n", url)
}
