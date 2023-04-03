package main

import (
	"fmt"
	"os"
)

// 写入数据到文件，文件存在的情况
func WriteFile() {

	var content string = "要写入文件的数据"
	// 转成[]byte 类型
	var s []byte = []byte(content)
	// 文件路径
	filePath := "./upload/a.txt"
	// 写入文件，文件存在
	// 0666 以读取和写入的权限
	err := os.WriteFile(filePath, s, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write success")
}

// 写入数据到文件，文件路径不存在的情况
// 使用perm 创建文件并写入数据
func WriteFilePathNotExist() {
	content := "文件路径不存在，创建文件，并写入数据"
	filePath := "./upload/b.txt"
	// 写入数据到文件，路径不存在
	// 以0666 可读可写的权限创建文件
	err := os.WriteFile(filePath, []byte(content), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write success")
}
func main() {
	WriteFilePathNotExist()
}
