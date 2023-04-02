package main

import (
	"fmt"
	"os"
)

// 创建目录
func main() {
	dirPath := "./upload1"
	// 创建目录，第一个参数为目录权限
	//err := os.Mkdir(dirPath, 0777)
	// 使用FileMode 常量定义权限
	err := os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("目录创建成功")

}
